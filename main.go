package main

import (
	_ "cleverbamboo.com/bee-shop-b2c/routers"
	_ "cleverbamboo.com/bee-shop-b2c/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
