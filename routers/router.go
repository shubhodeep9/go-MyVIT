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
		beego.NSNamespace("/menu",
			beego.NSInclude(
				// controllers/spotlight.go
				&controllers.MenuController{},
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
				// controllers/personalDetails.go
				&controllers.PersonalDetailsController{},
			),
		),

		beego.NSNamespace("/calCourses",
			beego.NSInclude(
				//controllers/familyDetails.go
				&controllers.CalCoursesController{},
			),
		),

		beego.NSNamespace("/educationalDetails",
			beego.NSInclude(
				//controllers/educationalDetails.go
				&controllers.EducationalDetailsController{},
			),
		),
		beego.NSNamespace("/familyDetails",
			beego.NSInclude(
				//controllers/familyDetails.go
				&controllers.FamilyDetailsController{},
			),
		),

		beego.NSNamespace("/hostelDetails",
			beego.NSInclude(
				//controllers/hostelDetails.go
				&controllers.HostelDetailsController{},
			),
		),
		beego.NSNamespace("/examSchedule",
			beego.NSInclude(
				// controllers/examSchedule.go
				&controllers.ExamSchedVtopBetaController{},
			),
		),
		beego.NSNamespace("/facphotos",
			beego.NSInclude(
				// controllers/facphotos.go
				&controllers.FacultyPhotosController{},
			),
		),
		beego.NSNamespace("/timetable2",
			beego.NSInclude(
				&controllers.Timetable2Controller{},
			),
		),
		beego.NSNamespace("/timetable",
			beego.NSInclude(
				&controllers.TimetableController{},
			),
		),

		beego.NSNamespace("/leaveRequest",
			beego.NSInclude(
				// controllers/examSchedule.go
				&controllers.LeaveRequestController{},
			),
		),
		beego.NSNamespace("/roomAllot",
			beego.NSInclude(
				// controllers/examSchedule.go
				&controllers.RoomAllotController{},
			),
		),
		beego.NSNamespace("/attendance",
			beego.NSInclude(
				// controllers/examSchedule.go
				&controllers.AttendanceController{},
			),
		),

		beego.NSNamespace("/attendanceDet",
			beego.NSInclude(
				// controllers/examSchedule.go
				&controllers.AttendanceDetailsController{},
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
