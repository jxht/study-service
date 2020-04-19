package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"os"
	"study-service/app/routes"
)

func main() {
	app := iris.New()
	routes.RegisterRoutes(app)
	err := app.Run(iris.Addr(":8888"))
	if err != nil {
		fmt.Printf("run server failed. err=%v\n", err)
		os.Exit(1)
	}
}
