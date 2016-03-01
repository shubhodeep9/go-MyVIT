package controllers

import (
	"github.com/astaxie/beego"
	"go-MyVIT/api"
)

// Operations about login
type AdvisorController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200
// @Failure 403 parameters missing
// @router / [get]
func (o *AdvisorController) Get() {
	regNo := o.Input().Get("regNo")
	psswd := o.Input().Get("psswd")
	if regNo != "" && psswd != "" {
		resp := api.Advisor(regNo, psswd)
		o.Data["json"] = resp
	}
	o.ServeJSON()
}
