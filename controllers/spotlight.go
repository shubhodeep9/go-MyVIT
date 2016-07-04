/*
SpotlightController
*/

package controllers

import (
	"github.com/astaxie/beego"
	"go-MyVIT/api"
)

// Operations about login
type SpotlightController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200
// @Failure 403 parameters missing
// @router / [get]
func (o *SpotlightController) Get() {
	campus := o.Ctx.Input.Param(":campus")
	var baseuri string
	if campus == "vellore" {
		baseuri = "https://academics.vit.ac.in"
	} else {
		baseuri = "https://academicscc.vit.ac.in"
	}
	resp := api.Spotlight(baseuri)
	o.Data["json"] = resp
	o.ServeJSON()
}
