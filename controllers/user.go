package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"github.com/salamander-mh/SalamanderBlog/models"
)

type UserController struct {
	beego.Controller
}

// @router /login/ [post]
func (this *UserController) Login() {

}

func (this *UserController) Get() {
	uid, _ := strconv.Atoi(this.Ctx.Input.Param(":uid"))
	user := models.GetUser(uid)
	this.Data["user"] = user
	this.TplName = "user.tpl"
}
