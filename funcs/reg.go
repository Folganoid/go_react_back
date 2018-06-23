package funcs

import (
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
	"fmt"
	"../config"
	"../models"
	"math/rand"

)

func RegUser(ctx iris.Context) {

	setup := config.Setup()

	db, err := gorm.Open("mysql", setup["db_user"] + ":" + setup["db_pass"] + "@/" + setup["db_name"] + "?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	user := models.NewUser(0, ctx.FormValue("login"), ctx.FormValue("pass"), ctx.FormValue("email"), randToken())

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
