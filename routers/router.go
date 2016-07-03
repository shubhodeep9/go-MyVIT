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
	ns := beego.NewNamespace("/campus/:campus",
		beego.NSNamespace("/login",
			beego.NSInclude(
				// controllers/login.go
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/coursepage/:category",
			beego.NSInclude(
				// controllers/login.go
				&controllers.CoursePageController{},
			),
		),
		beego.NSNamespace("/refresh",
			beego.NSInclude(
				&controllers.RefreshController{},
			),
		),
		beego.NSNamespace("/spotlight",
			beego.NSInclude(
				// controllers/spotlight.go
				&controllers.SpotlightController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
