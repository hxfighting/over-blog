package backend

import (
	"github.com/kataras/iris"
	"log"
)

func Login(ctx iris.Context) {
	log.Println(ctx.FormValues())
}

func Logout(ctx iris.Context) {

}
