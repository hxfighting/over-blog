package main

import (
	"blog/routes"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	routes.RegisterRoutes(app)
	if run := app.Run(iris.Addr(":8080")); run != nil {
		panic(run)
	}
}
