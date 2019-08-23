package controllers

import (
	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (v *VideoController) Get() {
	v.Ctx.WriteString("hello beego")
}
