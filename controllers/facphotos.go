package controllers

import (
	//"encoding/json"
	"github.com/astaxie/beego"
	"go-MyVIT/api"
	//"time"
)

type FacultyPhotosController struct {
	beego.Controller
}

// @Title Get faculty information
// @Description find object by objectid
// @Success 200
// @Failure 403 parameters missing
// @router / [post]
func (o *FacultyPhotosController) Post() {

	regNo := o.Input().Get("regNo")
	psswd := o.Input().Get("psswd")
	campus := o.Ctx.Input.Param(":campus")
	query := o.Input().Get("empid")

	var baseuri string
	if campus == "vellore" {
		baseuri = "https://vtop.vit.ac.in"
	} else {
		baseuri = "https://academicscc.vit.ac.in"
	}
	if regNo != "" && psswd != "" {
		o.Data["json"] = api.FacPhotos(regNo, psswd, query, baseuri)
	}
	o.ServeJSON()
}
