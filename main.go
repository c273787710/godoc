package main

import (
	_ "godoc/routers"
	_ "godoc/init"
	"github.com/astaxie/beego"

)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
