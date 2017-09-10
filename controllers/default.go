package controllers

import (
	"github.com/astaxie/beego"
	// "log"
	"markdown_blog_go/markdown"
	"net/url"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	name := c.GetString(":name")
	name, _ = url.QueryUnescape(name)
	if len(name) == 0 {
		file_infos := markdown.Search("")
		name = markdown.RemoveSuffix(file_infos[0].Name)
		c.Redirect("/"+name, 302)
	}
	c.Data["title"] = name
	c.Data["author"] = "bigzhu"
	c.Data["author_link"] = "http://bigzhu.lorstone.com/bigzhu"
	modify_time, err := markdown.GetFileModTime(name)
	if err != nil {
		c.Abort("404")
	}
	c.Data["modify_time"] = modify_time
	content, toc := markdown.GetContent(name)
	c.Data["toc"] = toc
	c.Data["content"] = content
	pre, old := markdown.PreAndOld(name)
	c.Data["pre"] = pre
	c.Data["old"] = old
	c.Data["quote_pre"] = url.QueryEscape(pre)
	c.Data["quote_old"] = url.QueryEscape(old)

	c.TplName = "index.tpl"
}
