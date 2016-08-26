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
	"strconv"
	"os"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.HttpPort = port
	}
	beego.Run()
}
