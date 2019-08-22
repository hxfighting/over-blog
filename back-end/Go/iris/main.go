package main

import (
	"blog/queue"
	"blog/routes"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	routes.RegisterRoutes(app)
	go func() {
		queue.HandleQueue()
	}()
	if run := app.Run(iris.Addr(":8080")); run != nil {
		panic(run)
	}
}
