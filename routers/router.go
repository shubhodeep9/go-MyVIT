// @APIVersion 1.0.0
// @Title MyVIT academics API
// @Description Simple scraping api written in Go for MyVIT app (Google Developers Group VIT Vellore)
// @Contact shubhodeep9@gmail.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go-MyVIT/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/login",
			beego.NSInclude(
				// controllers/login.go
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/timetable",
			beego.NSInclude(
				// controllers/timetable.go
				&controllers.TimetableController{},
			),
		),
		beego.NSNamespace("/facadvdet",
			beego.NSInclude(
				// controllers/advisor.go
				&controllers.AdvisorController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
