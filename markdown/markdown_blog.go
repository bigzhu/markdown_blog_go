package main

import (
	"fmt"
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

type File struct {
	name string
	time time.Time
}

type Files []File

func (a Files) Less(i, j int) bool { return a[j].time.Before(a[i].time) }
func (a Files) Len() int           { return len(a) }
func (a Files) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func getContent(name string) string {
	file, err := ioutil.ReadFile(path + name + ".md")
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	s := string(file)

	fmt.Printf(s)
	return s
	// content = name_file.read()
	// name_file.close()
	// if 'status: draft' in content:
	//     return '# 这是一个机密文件, 不允许查看!'
	// return content

}

func preAndOld(name string) (string, string) { //返回前一个文章后一个文章的名字
	file_infos := search("")
	file_name := name + ".md"
	s_len := len(file_infos)
	for index, value := range file_infos {
		if value.name == file_name {
			if s_len == 1 {
				return "", ""
			}
			if index == 1 { // 第一个
				return "", removeSuffix(file_infos[1].name)
			}
			if index == s_len-1 { // 最后一个
				return removeSuffix(file_infos[index-1].name), ""
			}
			return removeSuffix(file_infos[index-1].name), removeSuffix(file_infos[index+1].name)
		}
	}

	return "", ""
}
func removeSuffix(name string) string { // 删后缀
	var extension = filepath.Ext(name)
	name = name[0 : len(name)-len(extension)]
	return name
}
func search(search_name string) []File {
	// 获取所有文件
	var file_infos []File
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			name := file.Name()
			if !strings.Contains(name, search_name) {
				continue
			}
			// 取文件的时间
			fi, err := os.Stat(path + name)
			if err != nil {
				log.Fatal(err)
			}
			mtime := fi.ModTime()
			file_info := File{name, mtime}
			file_infos = append(file_infos, file_info)
			sort.Sort(Files(file_infos))
		}
	}
	return file_infos
}

// const MD_PATH = home + '/Dropbox/blog/'
func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return
	}
	var home = usr.HomeDir
	path = home + "/Dropbox/blog/"
	getContent("健身计划")
	// pre, old := preAndOld("健身计划")
	// fmt.Printf("%s\n", pre)
	// fmt.Printf("%s\n", old)
	// file_infos := search("bigzhu")
	// for index, value := range file_infos {
	// 	fmt.Printf("%s", index)
	// 	fmt.Printf(removeSuffix(value.name))
	// 	fmt.Printf(value.time.String())
	// }
}
