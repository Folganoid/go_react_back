package funcs

import (
	"github.com/kataras/iris"
)

func Routing(app *iris.Application) {

	app.Post("/reg", func(ctx iris.Context) {
		//registration
		RegUser(ctx)
	})

	app.Post("/user", func(ctx iris.Context) {
		//auth
		FetchUser(ctx)
	})

	app.Post("/profupdate", func(ctx iris.Context) {
		//update user
		UpdateUser(ctx)
	})

	app.Options("/user", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
	})
	app.Options("/profupdate", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
	})
	app.Options("/reg", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
	})
}
