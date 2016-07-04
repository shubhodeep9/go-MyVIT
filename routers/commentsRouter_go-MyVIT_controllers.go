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

	beego.GlobalControllerRouter["go-MyVIT/controllers:LoginController"] = append(beego.GlobalControllerRouter["go-MyVIT/controllers:LoginController"],
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

}
