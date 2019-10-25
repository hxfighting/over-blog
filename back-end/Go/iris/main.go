package main

import (
	"blog/queue"
	"blog/routes"
	stdContext "context"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Authorization"},
	}))
	app.AllowMethods(iris.MethodOptions)
	app.HandleDir("/css", "./public/css")
	app.HandleDir("/js", "./public/js")
	app.HandleDir("/image", "./public/image")
	app.HandleDir("/static", "./public/static")
	//app.RegisterView(iris.HTML("./views", ".html"))
	routes.RegisterRoutes(app)
	go func() {
		queue.HandleQueue()
	}()
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			// kill -SIGINT XXXX 或 Ctrl+c
			os.Interrupt,
			syscall.SIGINT, // register that too, it should be ok
			// os.Kill等同于syscall.Kill
			os.Kill,
			syscall.SIGKILL, // register that too, it should be ok
			// kill -SIGTERM XXXX
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			println("\n服务器关闭...")
			timeout := 5 * time.Second
			ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
			defer cancel()
			app.Shutdown(ctx)
		}
	}()
	if run := app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler, iris.WithOptimizations); run != nil {
		log.Fatalln(run.Error())
	}
}
