package fetch

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/kataras/iris"
	"./models"
)

func fetchUser(ctx *iris.Context) {
	db, err := gorm.Open("mysql", "t430:56195619@/go?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	user := models.User{}
	db.Where("login = ? AND pass= ?", ctx.FormValue("login"), ctx.FormValue("pass")).Find(&user)

	form := ctx.FormValues()

	fmt.Println(form)

	if user.Id > 0 {
		fmt.Println(ctx.JSON(user))
	} else {
		fmt.Println(ctx.JSON(user))
	}
}
