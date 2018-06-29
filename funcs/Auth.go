package funcs

import (
	"github.com/kataras/iris"
	"fmt"
	"../models"
	"time"
	"math/rand"
	"crypto/md5"
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
		ctx.WriteString("SYSTEM Error unauthorized 401")
	}
}

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
		ctx.WriteString("SYSTEM Error Access denied 401")
	}
}

/**
Check user
 */
func CheckUser(ctx *iris.Context, userid string, token string) bool {
	db := ConnectDB(*ctx)
	defer db.Close()

	user := models.User{}
	db.Where("id = ? AND token= ?", userid, token).Find(&user)

	if user.Id > 0 {
		return true
	} else {
		return false
	}
}

func RegUser(ctx iris.Context) {

	db := ConnectDB(ctx)
	defer db.Close()

	user := models.NewUser(0,
		ctx.FormValue("login"),
		ctx.FormValue("name"),
		ctx.FormValue("email"),
		time.Now().Unix(),
		0,
		GetMD5Hash(ctx.FormValue("pass")),
		ctx.FormValue("year"),
		randToken(),
	)

	db.Create(user)
	if db.Error != nil {
		fmt.Println(db.Error)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(db.Error.Error())
	}
}

/**
* create random token
 */
func randToken() string {
	b := make([]byte, 24)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

/**
* generate md5 string
 */
func GetMD5Hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
