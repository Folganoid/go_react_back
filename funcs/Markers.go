package funcs

import (
	"github.com/kataras/iris"
	"fmt"
	"../models"
)

func GetMarkers(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	db := ConnectDB(ctx)
	defer db.Close()

	markers := []models.Marker{}
	db.Where("userid = ?", ctx.FormValue("userid")).Find(&markers)

	fmt.Println(ctx.JSON(markers))
}

func GetForeignMarkers(ctx iris.Context) {

	db := ConnectDB(ctx)
	defer db.Close()

	user := models.User{}
	db.Where("login = ? AND allow_map = 1", ctx.FormValue("login")).Find(&user)

	if user.Id > 0 {
		markers := []models.Marker{}
		db.Where("userid = ?", user.Id).Find(&markers)
		fmt.Println(ctx.JSON(markers))
	} else {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error User not found or map is closed for watching!")
	}
}
