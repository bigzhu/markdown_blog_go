package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"sort"
	"time"
)

type File struct {
	name string
	time time.Time
}

type Files []File

func (a Files) Less(i, j int) bool { return a[j].time.Before(a[i].time) }
func (a Files) Len() int           { return len(a) }
func (a Files) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func search(path string, name string) []File {
	// 获取所有文件
	var file_infos []File
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			name := file.Name()
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
	var path = home + "/Dropbox/blog/"
	file_infos := search(path, "bigzhu")
	for index, value := range file_infos {
		fmt.Printf("%s", index)
		fmt.Printf(value.name)
		fmt.Printf(value.time.String())
	}
}
