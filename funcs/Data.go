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

/**
 * [Bike description]
 * @param {[type]} ctx iris.Context [description]
 */
func Bike(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	db := ConnectDB(ctx)
	defer db.Close()

	// POST
	if ctx.Method() == "POST" {

		bikes := []models.Bike{}
		db.Where("userid = ?", ctx.FormValue("userid")).Find(&bikes)
		fmt.Println(ctx.JSON(bikes))
	}

	// PUT
	if ctx.Method() == "PUT" {
		if len(ctx.FormValue("bike")) > 0 {

			userId, _ := strconv.ParseUint(ctx.FormValue("userid"), 10, 64)
			bike := models.NewBike(0, ctx.FormValue("bike"), userId)
			db.Create(&bike)

			fmt.Println(db.NewRecord(bike))
		}
	}

	// DELETE
	if ctx.Method() == "DELETE" {

		id, _ := strconv.ParseUint(ctx.FormValue("id"), 10, 64)

		bike := models.NewBike(id, "", 0)
		db.Delete(&bike)
	}
}

/**
 * [Tire description]
 * @param {[type]} ctx iris.Context [description]
 */
func Tire(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	db := ConnectDB(ctx)
	defer db.Close()

	// PUT
	if ctx.Method() == "PUT" {
		if len(ctx.FormValue("tire")) > 0 {

			userId, _ := strconv.ParseUint(ctx.FormValue("userid"), 10, 64)
			tire := models.NewTire(0, ctx.FormValue("tire"), userId)
			db.Create(&tire)

			fmt.Println(db.NewRecord(tire))
		}
	}

	// POST
	if ctx.Method() == "POST" {
		tires := []models.Tire{}
		db.Where("userid = ?", ctx.FormValue("userid")).Find(&tires)
		fmt.Println(ctx.JSON(tires))
	}

	// DELETE
	if ctx.Method() == "DELETE" {

		id, _ := strconv.ParseUint(ctx.FormValue("id"), 10, 64)

		tire := models.NewTire(id, "", 0)
		db.Delete(&tire)
	}

}

/**
 * [YearDist description]
 * @param {[type]} ctx iris.Context [description]
 */
func YearDist(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	db := ConnectDB(ctx)
	defer db.Close()

	// POST
	if ctx.Method() == "POST" {

		yearList := []models.YearDist{}
		db.Where("userid = ?", ctx.FormValue("userid")).Order("year desc", true).Find(&yearList)

		fmt.Println(ctx.JSON(yearList))
	}

	// PUT
	if ctx.Method() == "PUT" {

		userId, _ := strconv.ParseUint(ctx.FormValue("userid"), 10, 64)
		year, _ := strconv.ParseUint(ctx.FormValue("year"), 10, 64)
		dist, _ := strconv.ParseFloat(ctx.FormValue("dist"), 64)

		yearDist := models.NewYearDist(0, userId, year, ctx.FormValue("bike"), dist)
		db.Create(&yearDist)

		//fmt.Println(db.NewRecord(yearDist))
	}

	// DELETE
	if ctx.Method() == "DELETE" {

		id, _ := strconv.ParseUint(ctx.FormValue("id"), 10, 64)
		yearDist := models.NewYearDist(id, 0, 0, "", 0)
		db.Delete(&yearDist)
	}

}
