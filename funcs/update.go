package funcs

import (
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
	"fmt"
	"../config"
	"../models"
)

/**
User update
 */
func UpdateUser(ctx iris.Context) {

	setup := config.Setup()

	db, err := gorm.Open("mysql", setup["db_user"] + ":" + setup["db_pass"] + "@/" + setup["db_name"] + "?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	user := models.User{}
	db.Where("id = ? AND pass = ?", ctx.FormValue("userId"), ctx.FormValue("pass_old")).Find(&user)

	if user.Id > 0 {
		user.Token = randToken()
		user.Login = ctx.FormValue("login")
		user.Email = ctx.FormValue("email")
		user.Year = ctx.FormValue("year")
		user.Pass = GetMD5Hash(ctx.FormValue("pass"))

		db.Save(&user)
		user.Pass = "updated" // clear pass
		fmt.Println(ctx.JSON(user))
	} else {
		user.Pass = "false" // clear pass
		fmt.Println(ctx.JSON(user))
	}
}
