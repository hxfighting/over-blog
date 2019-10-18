package main

import (
	"blog/queue"
	"blog/routes"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Authorization"},
	}))
	app.AllowMethods(iris.MethodOptions)
	app.RegisterView(iris.HTML("./views", ".html"))
	routes.RegisterRoutes(app)
	go func() {
		queue.HandleQueue()
	}()
	if run := app.Run(iris.Addr(":8080")); run != nil {
		panic(run)
	}
}
