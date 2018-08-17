package funcs

import (
	"../models"
	"fmt"
	"github.com/kataras/iris"
	"strconv"
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

func GetBikeList(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	db := ConnectDB(ctx)
	defer db.Close()

	bikes := []models.Bike{}
	db.Where("userid = ?", ctx.FormValue("userid")).Find(&bikes)

	fmt.Println(ctx.JSON(bikes))
}

func GetTireList(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	db := ConnectDB(ctx)
	defer db.Close()

	tires := []models.Tire{}
	db.Where("userid = ?", ctx.FormValue("userid")).Find(&tires)

	fmt.Println(ctx.JSON(tires))
}

func AddBike(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	if len(ctx.FormValue("bike")) > 0 {

		db := ConnectDB(ctx)
		defer db.Close()

		userId, _ := strconv.ParseUint(ctx.FormValue("userid"), 10, 64)
		bike := models.NewBike(0, ctx.FormValue("bike"), userId)
		db.Create(&bike)

		fmt.Println(db.NewRecord(bike))
	}
}

func AddTire(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	if len(ctx.FormValue("tire")) > 0 {

		db := ConnectDB(ctx)
		defer db.Close()

		userId, _ := strconv.ParseUint(ctx.FormValue("userid"), 10, 64)
		tire := models.NewTire(0, ctx.FormValue("tire"), userId)
		db.Create(&tire)

		fmt.Println(db.NewRecord(tire))
	}
}
