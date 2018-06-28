package funcs

import (

	"github.com/kataras/iris"
	"fmt"
	"../models"
	"math/rand"
	"time"
	"crypto/md5"
)

/**
User registration
 */
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
