package routes

import "github.com/kataras/iris/v12"

func RegisterRoutes(app *iris.Application) {
	//默认页面
	app.Get("/", func(ctx iris.Context) {
		_, _ = ctx.WriteString("welcome to service")
	})
	//404
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		//m.ServeHTTP(ctx)
		_, _ = ctx.WriteString("404 Not Found")
	})
}
