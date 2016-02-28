package main

import (
	_ "go-MyVIT/docs"
	_ "go-MyVIT/routers"

	"go-MyVIT/Godeps/_workspace/src/github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
