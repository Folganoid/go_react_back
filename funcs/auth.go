package funcs

import (
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
	"fmt"
	"../models"
	"../config"
)

/**
Get user
 */
func FetchUser(ctx iris.Context) {

	setup := config.Setup()

	db, err := gorm.Open("mysql", setup["db_user"] + ":" + setup["db_pass"] + "@/" + setup["db_name"] + "?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	user := models.User{}
	db.Where("login = ? AND pass= ?", ctx.FormValue("login"), GetMD5Hash(ctx.FormValue("pass"))).Find(&user)

	if user.Id > 0 {
		user.Pass = "" // clear pass
		fmt.Println(ctx.JSON(user))
	} else {
		fmt.Println(ctx.JSON(user))
	}
}
