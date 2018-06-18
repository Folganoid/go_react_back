package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/iris-contrib/middleware/cors"
	"./funcs"
)


func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.Default())

	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Post("/user", func(ctx iris.Context) {

		//auth
		funcs.FetchUser(ctx)

	})
	app.Options("/user", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
	})

	app.Run(iris.Addr(":3001"), iris.WithoutServerError(iris.ErrServerClosed))
}
