// @APIVersion 1.0.0
// @Title MyVIT academics API
// @Description Simple scraping api written in Go for MyVIT app (Google Developers Group VIT Vellore)
// @Contact shubhodeep9@gmail.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"go-MyVIT/controllers"
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
				// controllers/coursepage.go
				&controllers.CoursePageController{},
			),
		),
		beego.NSNamespace("/refresh",
			beego.NSInclude(
				// controllers/refresh.go
				&controllers.RefreshController{},
			),
		),
		beego.NSNamespace("/spotlight",
			beego.NSInclude(
				// controllers/spotlight.go
				&controllers.SpotlightController{},
			),
		),
		beego.NSNamespace("/pic",
			beego.NSInclude(
				// controllers/pic.go
				&controllers.PicController{},
			),
		),
		beego.NSNamespace("/facdet",
			beego.NSInclude(
				&controllers.FacultyController{},
			),
		),
		beego.NSNamespace("/messages",
			beego.NSInclude(
				// controllers/messages.go
				&controllers.MessageController{},
			),
		),
		beego.NSNamespace("/personalDetails",
			beego.NSInclude(
				// controllers/messages.go
				&controllers.PersonalDetailsController{},
			),
		),
	)
	stats := beego.NewNamespace("/admin",
		beego.NSNamespace("/stats",
			beego.NSInclude(
				// controllers/stats.go
				&controllers.StatsController{},
			),
		),
		beego.NSNamespace("/message/:function",
			beego.NSInclude(
				&controllers.GcmMessageController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(stats)
}
