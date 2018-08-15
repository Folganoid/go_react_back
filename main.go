package main

import (
	"./funcs"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func notFound(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
	//ctx.View("errors/404.html")
}

func internalServerError(ctx iris.Context) {
	ctx.WriteString("Oups something went wrong, try again")
}

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.Default())
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	// routing
	funcs.Routing(app)

	// starting server
	app.Run(iris.Addr(":3001"), iris.WithoutServerError(iris.ErrServerClosed))
}
