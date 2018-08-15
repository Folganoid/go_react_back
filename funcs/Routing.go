package funcs

import (
	"../config"
	"github.com/kataras/iris"
)

func Routing(app *iris.Application) {

	setup := config.Setup()

	app.Post("/user", func(ctx iris.Context) { FetchUser(ctx) })
	app.Post("/profupdate", func(ctx iris.Context) { UpdateUser(ctx) })
	app.Post("/reg", func(ctx iris.Context) { RegUser(ctx) })
	app.Post("/get_markers", func(ctx iris.Context) { GetMarkers(ctx) })
	app.Post("/get_foreign_markers", func(ctx iris.Context) { GetForeignMarkers(ctx) })
	app.Post("/get_year_data", func(ctx iris.Context) { GetYear(ctx) })
	app.Post("/get_stat_data", func(ctx iris.Context) { GetStat(ctx) })
	app.Post("/get_bike_list", func(ctx iris.Context) { GetBikeList(ctx) })
	app.Post("/get_tire_list", func(ctx iris.Context) { GetTireList(ctx) })

	app.Options("/user", func(ctx iris.Context) { ctx.Header("Access-Control-Allow-Origin", setup["frontPath"]) })
	app.Options("/profupdate", func(ctx iris.Context) { ctx.Header("Access-Control-Allow-Origin", setup["frontPath"]) })
	app.Options("/reg", func(ctx iris.Context) { ctx.Header("Access-Control-Allow-Origin", setup["frontPath"]) })
	app.Options("/check_user", func(ctx iris.Context) { ctx.Header("Access-Control-Allow-Origin", setup["frontPath"]) })
	app.Options("/get_markers", func(ctx iris.Context) { ctx.Header("Access-Control-Allow-Origin", setup["frontPath"]) })
	app.Options("/get_foreign_markers", func(ctx iris.Context) { ctx.Header("Access-Control-Allow-Origin", setup["frontPath"]) })
	app.Options("/get_year_data", func(ctx iris.Context) { ctx.Header("Access-Control-Allow-Origin", setup["frontPath"]) })
	app.Options("/get_stat_data", func(ctx iris.Context) { ctx.Header("Access-Control-Allow-Origin", setup["frontPath"]) })
	app.Options("/get_bike_list", func(ctx iris.Context) { ctx.Header("Access-Control-Allow-Origin", setup["frontPath"]) })
	app.Options("/get_tire_list", func(ctx iris.Context) { ctx.Header("Access-Control-Allow-Origin", setup["frontPath"]) })
}
