package funcs

import (
	"github.com/kataras/iris"
	"fmt"
	"../models"
)

/**
User update
 */
func UpdateUser(ctx iris.Context) {

	db := ConnectDB(ctx)
	defer db.Close()

	user := models.User{}
	db.Where("id = ? AND pass = ?", ctx.FormValue("userId"), GetMD5Hash(ctx.FormValue("pass_old"))).Find(&user)

	if user.Id > 0 {
		user.Token = randToken()
		user.Login = ctx.FormValue("login")
		user.Email = ctx.FormValue("email")
		user.Year = ctx.FormValue("year")
		user.Pass = GetMD5Hash(ctx.FormValue("pass"))

		db.Save(&user)
		user.Pass = "" // clear pass
		fmt.Println(ctx.JSON(user))
	} else {
		ctx.StatusCode(401)
		ctx.WriteString("SYSerror Access denied 401")
	}
}
