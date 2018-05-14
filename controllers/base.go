package controllers

import (
	"github.com/astaxie/beego"
	"github.com/salamander-mh/SalamanderBlog/models"
)

type BaseController struct {
	beego.Controller
	User models.User
}

func (c *BaseController) Prepare() {

}