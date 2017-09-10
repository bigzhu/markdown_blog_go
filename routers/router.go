package routers

import (
	"github.com/astaxie/beego"
	"markdown_blog_go/controllers"
)

func init() {
	beego.Router("/?:name", &controllers.MainController{})
}
