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
type GcmMessageController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200
// @Failure 403 parameters missing
// @router / [get]
func (o *GcmMessageController) Get() {
	function := o.Ctx.Input.Param(":function")
	if function == "send" {
		message := o.Input().Get("message")
		o.Data["json"] = api.GcmSender(message)
	} else if function == "register" {
		reg := o.Input().Get("regID")
		o.Data["json"] = api.GcmRegister(reg)
		// } else if function == "upload" {

	}
	o.ServeJSON()
}
