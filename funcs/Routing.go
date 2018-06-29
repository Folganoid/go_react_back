package funcs

import (
	"github.com/kataras/iris"
)

func Routing(app *iris.Application) {

	app.Post("/user", func(ctx iris.Context) {FetchUser(ctx)})
	app.Post("/profupdate", func(ctx iris.Context) {UpdateUser(ctx)})
	app.Post("/reg", func(ctx iris.Context) {RegUser(ctx)	})
	app.Post("/get_markers", func(ctx iris.Context) {GetMarkers(ctx)})


	app.Options("/user", func(ctx iris.Context) {	ctx.Header("Access-Control-Allow-Origin", "*")})
	app.Options("/profupdate", func(ctx iris.Context) {ctx.Header("Access-Control-Allow-Origin", "*")})
	app.Options("/reg", func(ctx iris.Context) {ctx.Header("Access-Control-Allow-Origin", "*")})
	app.Options("/check_user", func(ctx iris.Context) {ctx.Header("Access-Control-Allow-Origin", "*")})
	app.Options("/get_markers", func(ctx iris.Context) {ctx.Header("Access-Control-Allow-Origin", "*")})
}
