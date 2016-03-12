package routers

import (
	"github.com/astaxie/beego"
)

func init() {

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

	beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}
