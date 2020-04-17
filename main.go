package main

import (
	"github.com/kataras/iris/v12"
	"study-service/app/routes"
)

func main() {
	app := iris.New()
	routes.RegisterRoutes(app)
	app.Run(iris.Addr(":8888"))
}
