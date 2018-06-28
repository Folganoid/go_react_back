package funcs

import (
	"github.com/kataras/iris"
	"fmt"
	"../models"
)

/**
Get user
 */
func FetchUser(ctx iris.Context) {

	db := ConnectDB(ctx)
	defer db.Close()

	user := models.User{}
	db.Where("login = ? AND pass= ?", ctx.FormValue("login"), GetMD5Hash(ctx.FormValue("pass"))).Find(&user)

	if user.Id > 0 {
		user.Token = randToken()
		db.Save(&user)
		user.Pass = "" // clear pass
		fmt.Println(ctx.JSON(user))
	} else {
		ctx.StatusCode(401)
		ctx.WriteString("SYSerror unauthorized 401")
	}
}
