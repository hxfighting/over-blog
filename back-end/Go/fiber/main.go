package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"

	"github.com/ohdata/blog/configs"
	initialize "github.com/ohdata/blog/init"
	"github.com/ohdata/blog/internal/handlers"
	"github.com/ohdata/blog/internal/middlewares"
	"github.com/ohdata/blog/tools"
	"github.com/ohdata/blog/tools/json"
)

func main() {
	initialize.Setup()
	defer initialize.Close()
	c := configs.Config.Server
	cfg := fiber.Config{
		ServerHeader:          c.Name,
		CaseSensitive:         true,
		Immutable:             true,
		BodyLimit:             c.MaxBody,
		Concurrency:           c.Concurrency,
		ReadTimeout:           c.ReadTimeout,
		WriteTimeout:          c.WriteTimeout,
		DisableStartupMessage: !c.Debug,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err == fiber.ErrRequestTimeout {
				return nil
			}
			return tools.Error(ctx, http.StatusText(http.StatusInternalServerError))
		},
		JSONEncoder: json.Marshal,
	}
	r := fiber.New(cfg)

	middlewares.RegisterMiddleware(r)
	handlers.RegisterRoute(r)

	// 404
	r.Use(func(ctx *fiber.Ctx) error {
		return tools.Error(ctx, http.StatusText(http.StatusNotFound))
	})

	go func() {
		if err := r.Listen(c.Addr); err != nil && err != http.ErrServerClosed {
			log.Fatalln("listen error: ", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")
	if err := r.Shutdown(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Server exits.")
}
