/**********************************************
** @Des: 入口文件
** @Author: haodaquan
** @Date:   2017-11-02 02
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-11-05 10:19:08
***********************************************/
package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/george518/PPGo_ClipImg/coreImg"
	"github.com/george518/PPGo_ClipImg/lib"
)

func main() {
	// go run main.go -f=./example/images/image.jpg -p=center -h=300 -w=200 -m=10
	var width, height, margin int
	var path, postion string

	flag.StringVar(&path, "f", "", "图片地址或者图片文件夹地址")
	flag.StringVar(&postion, "p", "center", "top|bottom|left|right|center")
	flag.IntVar(&height, "h", 300, "新图片高度")
	flag.IntVar(&width, "w", 300, "新图片宽度")
	flag.IntVar(&margin, "m", 20, "边距")

	flag.Parse()

	//判断入参规则
	if path == "" {
		fmt.Println("file is required")
		return
	}

	if strings.Contains("top|bottom|left|right|center", postion) == false {
		fmt.Println("illegal posion: top|bottom|left|right|center")
		return
	}

	if _, err := os.Stat(path); err != nil {
		fmt.Println("file or dir not exists")
		return
	}

	//打开图片
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	bt := bytes.NewBuffer(file)
	m, _, err := image.Decode(bt)
	if err != nil {
		log.Fatal(err)
	}

	//截取核心图片
	x0, x1, y0, y1 := coreImg.CoreImg(m)

	//计算核心图片大小
	dx1 := x1 - x0
	dy1 := y1 - y0

	var rate_width int //核心图片缩小或扩大的宽度
	var rate_height int
	img_width := width - 2*margin //去掉边距
	img_height := height - 2*margin

	// fmt.Println("缩略图宽高", img_width, img_height)
	// fmt.Println("核心图宽高", dx1, dy1)

	if dy1 > dx1 {
		rate_height = img_height
		rate_width = rate_height * dx1 / dy1
		//如果比例还比底图大
		if rate_width > img_width {
			rate_width = img_width
			rate_height = rate_width * dy1 / dx1
		}
	} else {
		rate_width = img_width
		rate_height = rate_width * dy1 / dx1
		if rate_height > img_height {
			rate_height = img_height
			rate_width = rate_height * dx1 / dy1
		}
	}
	// fmt.Println("核心图坐标：", x0, y0, x1, y1)
	// fmt.Println("缩略图大小：", rate_width, rate_height)
	//截取图片
	corePic := imaging.Crop(m, image.Rect(x0, y0, x1, y1))
	// dst := imaging.Fit(corePic, rate_width, rate_height, imaging.Lanczos)
	dst := imaging.Fill(corePic, rate_width, rate_height, imaging.Center, imaging.Lanczos)

	//创建地图并计算覆盖位置
	px, py := lib.Posion(rate_width, rate_height, width, height, postion)

	fmt.Println(rate_width, rate_height, width, height, postion)
	fmt.Println(px, py)
	bgdst := image.NewRGBA(image.Rect(0, 0, width, height))
	// white := color.RGBA{255, 0, 255, 255}
	white := m.At(0, 0)
	draw.Draw(bgdst, bgdst.Bounds(), &image.Uniform{white}, image.Pt(0, 0), draw.Over)
	draw.Draw(bgdst, bgdst.Bounds(), dst, dst.Bounds().Min.Sub(image.Pt(px, py)), draw.Over)

	//创建地址和文件名，并保存图片
	newPath := lib.MakeDir(path)
	err = imaging.Save(bgdst, newPath)
	if err != nil {
		log.Fatal(err)
	}

}
