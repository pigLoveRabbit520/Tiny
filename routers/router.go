package routers

import (
	"github.com/astaxie/beego"
	"github.com/salamander-mh/SalamanderBlog/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/user/:uid([0-9]+", &controllers.UserController{})
}