/*
@Author Shubhodeep Mukherjee
@Organization Google Developers Group VIT Vellore
	Isn't Go sexy?
	I know right!!
	Just like Shubhodeep
	I mean, have you seen the guy? xP
	#GDGSwag
*/

package controllers

import (
	"github.com/astaxie/beego"
	"go-MyVIT/api"
	"strings"
)

// Operations about login
type CoursePageController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200
// @Failure 403 parameters missing
// @router / [post]
func (o *CoursePageController) Post() {
	regNo := o.Input().Get("regNo")
	psswd := o.Input().Get("psswd")
	campus := o.Ctx.Input.Param(":campus")
	var baseuri string
	if campus == "vellore" {
		baseuri = "https://vtop.vit.ac.in"
	} else {
		baseuri = "https://academicscc.vit.ac.in"
	}
	if regNo != "" && psswd != "" {
		category := o.Ctx.Input.Param(":category")
		if category == "courses" {
			o.Data["json"] = api.CourseCoursesPage(regNo, psswd, baseuri)
		} else if category == "slots" {
			coursekey := o.Input().Get("crs")
			o.Data["json"] = api.CourseSlotsPage(regNo, psswd, baseuri, coursekey)
		} else if category == "faculties" {
			coursekey := o.Input().Get("crs")
			slt := strings.Replace(o.Input().Get("slt"), " ", "+", 1)
			o.Data["json"] = api.CourseFacPage(regNo, psswd, baseuri, coursekey, slt)
		} else if category == "data" {
			coursekey := o.Input().Get("crs")
			slt := strings.Replace(o.Input().Get("slt"), " ", "+", 1)
			fac := o.Input().Get("fac")
			o.Data["json"] = api.CourseDataPage(regNo, psswd, baseuri, coursekey, slt, fac)
		} else {
			o.Abort("404")
		}
		o.Ctx.Output.Header("Cookie", api.CookieReturn(regNo))
	}
	o.ServeJSON()
}
