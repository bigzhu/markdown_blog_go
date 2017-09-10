package controllers

import (
	"github.com/astaxie/beego"
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
	c.Data["modify_time"] = markdown.GetFileModTime(name)
	c.Data["toc"] = "will"
	c.Data["content"] = markdown.GetContent(name)
	pre, old := markdown.PreAndOld(name)
	c.Data["pre"] = pre
	c.Data["old"] = old
	c.Data["quote_pre"] = url.QueryEscape(pre)
	c.Data["quote_old"] = url.QueryEscape(old)

	c.TplName = "index.tpl"
}
