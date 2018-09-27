package funcs

import (
	"../models"
	"fmt"
	"github.com/kataras/iris"
	"regexp"
	"strconv"
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

func Marker(ctx iris.Context) {

	check := CheckUser(&ctx, ctx.FormValue("userid"), ctx.FormValue("token"))

	if !check {
		ctx.StatusCode(401)
		ctx.WriteString("SYSTEM Error Access denied 401")
		return
	}

	db := ConnectDB(ctx)
	defer db.Close()

	userId, _ := strconv.ParseUint(ctx.FormValue("userid"), 10, 64)
	name := ctx.FormValue("name")
	subname := ctx.FormValue("subname")
	link := ctx.FormValue("link")
	coord := ctx.FormValue("coord")
	color := ctx.FormValue("color")

	match, _ := regexp.MatchString("^\\d{1,3}\\.\\d+\\D+\\d{1,3}\\.\\d+$", coord)
	crd, _ := regexp.Compile("\\d{1,3}\\.\\d+")

	crds := crd.FindAllString(coord, -1)

	x, err := strconv.ParseFloat(crds[0], 64)
	y, err2 := strconv.ParseFloat(crds[1], 64)

	if !match || err != nil || err2 != nil {
		ctx.StatusCode(500)
		ctx.WriteString("Bad data")
		return
	}

	// POST
	if ctx.Method() == "POST" {
		if len(name) > 0 && len(color) > 0 && len(coord) > 0 {
			marker := models.NewMarker(0, userId, x, y, name, subname, link, color)
			db.Create(&marker)
		}
	}

	// PUT
	if ctx.Method() == "PUT" {

	}
}
