package markdown

import (
	"fmt"
	"net/url"
	"strings"
)

func FormatToc(title string) string { // 生成 toc 的 markdown 格式

	space := ""
	if strings.HasPrefix(title, "###") {
		space = "  "
	} else if strings.HasPrefix(title, "##") {
		space = " "
	}
	title = strings.Replace(title, "#", "", -1)
	escaped_title := url.QueryEscape(title)
	return fmt.Sprintf("%s* [%s](#%s)\n", space, title, escaped_title)
}

func CreateToc(content string) string { //取出 toc
	contents := strings.Split(content, "\n")
	var toc string

	for _, line := range contents {

		if strings.HasPrefix(line, "#") {
			toc += FormatToc(line)
		}
	}
	return toc
}
