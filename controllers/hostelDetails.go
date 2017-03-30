/*
hostel details controller
*/

package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"go-MyVIT/api"
)

// Operations about login
type HostelDetailsController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200
// @Failure 403 parameters missing
// @router / [post]
func (o *HostelDetailsController) Post() {
	//fmt.Println("controller")
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
		resp := api.ShowHostelDetails(regNo, psswd, baseuri)
		o.Data["json"] = resp
	}
	o.ServeJSON()
}