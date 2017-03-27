package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["go-MyVIT/controllers:CoursePageController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:CoursePageController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:FacultyController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:FacultyController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:GcmMessageController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:GcmMessageController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:LoginController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:LoginController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:MessageController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:MessageController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:PersonalDetailsController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:PersonalDetailsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:PicController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:PicController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:RefreshController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:RefreshController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:SpotlightController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:SpotlightController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:StatsController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:StatsController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})
	beego.GlobalControllerRouter["go-MyVIT/controllers:ExamScheduleController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:ExamScheduleController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

}
