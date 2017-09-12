package routers

import (
	"github.com/astaxie/beego"
	"markdown_blog_go/controllers"
)

func init() {
	beego.Router("/clear", &controllers.ClearController{})
	beego.Router("/?:name", &controllers.MainController{})
	beego.Router("/blog/?:name", &controllers.BlogController{})
}
