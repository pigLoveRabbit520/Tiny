package routers

import (
	"github.com/astaxie/beego"
	"github.com/salamander-mh/SalamanderBlog/controllers"
)

func init() {
	ns := beego.NewNamespace("/admin",

		beego.NSNamespace("/contents",
			beego.NSInclude(
				&controllers.ContentController{},
			),
		),
	)
	beego.AddNamespace(ns)
}