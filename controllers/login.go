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
)

// Operations about login
type LoginController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200
// @Failure 403 parameters missing
// @router / [post]
func (o *LoginController) Post() {
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
		resp := api.LogIn(regNo, psswd, baseuri)
		o.Data["json"] = resp
	}
	o.Ctx.Output.Header("Cookie", api.CookieReturn(regNo))
	o.ServeJSON()
}
