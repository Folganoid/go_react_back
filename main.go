package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/iris-contrib/middleware/cors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id    uint64
	Login string
	Pass  string
	Email string
	Token string
}

type Result struct {
	Result int8
}

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

		db, err := gorm.Open("mysql", "fg:5619@/go?charset=utf8")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		var user User
		db.Where("login = ? AND pass= ?", ctx.FormValue("login"), ctx.FormValue("pass")).Find(&user)

		form := ctx.FormValues()

		fmt.Println(form)

		if user.Id > 0 {
			fmt.Println(ctx.JSON(User(user)))
		} else {
			fmt.Println(ctx.JSON(User{}))
		}

	})
	app.Options("/user", func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
	})

	app.Run(iris.Addr(":3001"), iris.WithoutServerError(iris.ErrServerClosed))
}
