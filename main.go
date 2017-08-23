/*
@Author Shubhodeep Mukherjee
@Organization Google Developers Group VIT Vellore
	Isn't Go sexy?
	I know right!!
	Just like Shubhodeep
	I mean, have you seen the guy? xP
	#GDGSwag
*/

package main

import (
	"github.com/astaxie/beego"
	_ "go-MyVIT/docs"
	_ "go-MyVIT/routers"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}
	if os.Getenv("DEVENV") == "remote" {
		beego.BConfig.Listen.EnableHTTPS = true
		beego.BConfig.Listen.HTTPSPort = 10443
		beego.BConfig.Listen.HTTPSCertFile = "/home/shubhodeep9/go/src/go-MyVIT/conf/fullchain.pem"
		beego.BConfig.Listen.HTTPSKeyFile = "/home/shubhodeep9/go/src/go-MyVIT/conf/privkey.pem"
	}
	beego.Run()
}
