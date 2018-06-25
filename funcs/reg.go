package funcs

import (
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
	"fmt"
	"../config"
	"../models"
	"math/rand"
	"time"
	"crypto/md5"
)

/**
User registration
 */
func RegUser(ctx iris.Context) {

	setup := config.Setup()

	db, err := gorm.Open("mysql", setup["db_user"] + ":" + setup["db_pass"] + "@/" + setup["db_name"] + "?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	user := models.NewUser(0,
		ctx.FormValue("login"),
		ctx.FormValue("name"),
		ctx.FormValue("email"),
		time.Now(),
		0,
		GetMD5Hash(ctx.FormValue("pass")),
		ctx.FormValue("year"),
		randToken(),
		)

	db.Create(user)
	if db.Error != nil {
		fmt.Println(err)
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
