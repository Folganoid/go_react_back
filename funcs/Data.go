package funcs

import (
	"github.com/kataras/iris"
	"fmt"
	"../models"
)

/**
Get year statistic by usaer Id
 */
func GetYear(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	db := ConnectDB(ctx)
	defer db.Close()

	years := []models.Year{}
	db.Where("userid = ?", ctx.FormValue("userid")).Find(&years)

	fmt.Println(ctx.JSON(years))
}

/**
Get statistic by user Id
 */
func GetStat(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	db := ConnectDB(ctx)
	defer db.Close()

	stats := []models.Stat{}
	db.Where("userid = ?", ctx.FormValue("userid")).Find(&stats)

	fmt.Println(ctx.JSON(stats))
}
