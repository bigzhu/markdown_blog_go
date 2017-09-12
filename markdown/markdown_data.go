package markdown

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var path string
var files_cache = make(map[string][]File)

type File struct {
	Name string
	Time time.Time
}

type Files []File

func (a Files) Less(i, j int) bool { return a[j].Time.Before(a[i].Time) }
func (a Files) Len() int           { return len(a) }
func (a Files) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func GetContent(name string) (string, string) {
	file, err := ioutil.ReadFile(path + name + ".md")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	s := []byte(file)
	if strings.Contains(string(s), "status: draft") {
		s = []byte("# 这是一个机密文件, 不允许查看!")
	}
	toc := CreateToc(string(s))
	// fmt.Printf(toc)
	unsafe := blackfriday.MarkdownCommon(s)
	html := string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

	b_toc := []byte(toc)
	toc = string(blackfriday.MarkdownCommon(b_toc)) // 不能再执行SanitizeBytes, 会把 anchor 中的 %3A 转回:, 导致匹配不上
	// fmt.Printf("%s", toc)
	// toc = string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))
	return html, toc
	// content = name_file.read()
	// name_file.close()
	// if 'status: draft' in content:
	//     return '# 这是一个机密文件, 不允许查看!'
	// return content

}

func PreAndOld(name string) (string, string) { //返回前一个文章后一个文章的名字
	file_infos := Search("")
	file_name := name + ".md"
	s_len := len(file_infos)
	for index, value := range file_infos {
		// fmt.Printf("index: %d", index)
		if value.Name == file_name {
			if s_len == 1 {
				return "", ""
			}
			if index == 0 { // 第一个
				return "", RemoveSuffix(file_infos[1].Name)
			}
			if index == s_len-1 { // 最后一个
				return RemoveSuffix(file_infos[index-1].Name), ""
			}
			return RemoveSuffix(file_infos[index-1].Name), RemoveSuffix(file_infos[index+1].Name)
		}
	}

	return "", ""
}
func RemoveSuffix(name string) string { // 删后缀
	var extension = filepath.Ext(name)
	name = name[0 : len(name)-len(extension)]
	return name
}
func GetFileModTime(name string) (time time.Time, err error) {
	fi, err := os.Stat(path + name + ".md")
	if err != nil {
		return
	}
	return fi.ModTime(), err
}
func Search(search_name string) []File {
	// 获取所有文件
	var file_infos []File

	value, ok := files_cache[search_name]
	if ok {
		return value
	} else {
		fmt.Printf("%v not in files_cache\n", search_name)
	}

	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			name := file.Name()
			if name == "search.md" {
				continue
			} else if strings.HasSuffix(name, ".md") {
			} else {
				continue
			}
			if !strings.Contains(name, search_name) {
				continue
			}
			// 取文件的时间
			fi, err := os.Stat(path + name)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			mtime := fi.ModTime()
			file_info := File{name, mtime}
			file_infos = append(file_infos, file_info)
			sort.Sort(Files(file_infos))
		}
	}
	fmt.Printf("%v\n", len(file_infos))
	files_cache[search_name] = file_infos
	return file_infos
}

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	var home = usr.HomeDir
	path = home + "/Dropbox/blog/"
	// GetContent("健身计划")
}

// const MD_PATH = home + '/Dropbox/blog/'
func main() {

	// pre, old := preAndOld("健身计划")
	// fmt.Printf("%s\n", pre)
	// fmt.Printf("%s\n", old)
	// file_infos := search("bigzhu")
	// for index, value := range file_infos {
	// 	fmt.Printf("%s", index)
	// 	fmt.Printf(RemoveSuffix(value.Name))
	// 	fmt.Printf(value.Time.String())
	// }
}
