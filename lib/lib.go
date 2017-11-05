/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-11-02 13:30:56
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-11-05 00:18:32
***********************************************/
package lib

import (
	"fmt"
	"os"
	"strings"
)

func MakeDir(file string) string {

	// fmt.Println(runtime.GOOS) //操作系统
	// fmt.Println(runtime.GOARCH)
	path := strings.Split(file, "/")
	var newFilePath = ""
	for k, p := range path {
		if k != len(path)-1 {
			newFilePath += p + "/"
		}
	}
	newFilePath += "CoreImages/"
	CreateDir(newFilePath)
	return newFilePath + path[len(path)-1]
}

//判断目录是否存在
func IsDir(name string) bool {
	fi, err := os.Stat(name)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}

	return fi.IsDir()
}

//创建多级目录
func CreateDir(name string) bool {
	if IsDir(name) {
		return true
	}

	if createDirImpl(name) {
		return true
	} else {
		return false
	}
}

func createDirImpl(name string) bool {
	err := os.MkdirAll(name, 0777)
	if err == nil {
		return true
	} else {
		return false
	}
}

//计算合图位置 支持上下左右居中
func Posion(x0, y0, x1, y1 int, ps string) (int, int) {
	x := x1 - x0
	y := y1 - y0
	switch ps {
	case "left":
		return 0, int(y / 2)
	case "right":
		return int(x), int(y / 2)
	case "top":
		return int(x / 2), 0
	case "bottom":
		return int(x / 2), int(y)
	case "center":
		return int(x / 2), int(y / 2)
	default:
		return int(x / 2), int(y / 2)
	}
}
