package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["go-MyVIT/controllers:AcademicsController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:AcademicsController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:AdvisorController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:AdvisorController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:AttendanceController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:AttendanceController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:CoursePageController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:CoursePageController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:ExamScheduleController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:ExamScheduleController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:LoginController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:LoginController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:TimetableController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:TimetableController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

}
