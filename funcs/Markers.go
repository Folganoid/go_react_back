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
