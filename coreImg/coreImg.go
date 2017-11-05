/**********************************************
** @Des: 获取核心图片坐标，必须是底色为白色或者单色的图片
** @Author: haodaquan
** @Date:   2017-11-02 15:31:37
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-11-04 23:56:11
***********************************************/
package coreImg

import (
	"image"
	"image/color"
	_ "image/jpeg"
)

const (
	line  = 1
	param = 20
)

func CoreImg(m image.Image) (x0, x1, y0, y1 int) {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	bgColor := m.At(0, 0)

	chMinX := make(chan int, 1)
	chMaxX := make(chan int, 1)
	chMinY := make(chan int, 1)
	chMaxY := make(chan int, 1)

	go leftMinX(m, bgColor, dx, dy, chMinX)
	go leftMaxX(m, bgColor, dx, dy, chMaxX)
	go topMinY(m, bgColor, dx, dy, chMinY)
	go topMaxY(m, bgColor, dx, dy, chMaxY)

	x0 = <-chMinX
	x1 = <-chMaxX
	y0 = <-chMinY
	y1 = <-chMaxY

	return
}

func leftMinX(m image.Image, bgColor color.Color, dx, dy int, chMinX chan int) {
	for i := 0; i < dx; i = i + line {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			if isDiffColor(bgColor, colorRgb) {
				chMinX <- i - line + 1
				return
			}
		}
	}
	chMinX <- dx - 1
	return
}

func leftMaxX(m image.Image, bgColor color.Color, dx, dy int, chMaxX chan int) {
	for i := dx; i > 0; i = i - line {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i-1, j)
			if isDiffColor(bgColor, colorRgb) {
				chMaxX <- i - 1 + line
				return
			}
		}
	}
	chMaxX <- dx - 1
	return
}

func topMaxY(m image.Image, bgColor color.Color, dx, dy int, chMaxY chan int) {
	for i := dy; i > 0; i = i - line {
		for j := 0; j < dx; j++ {
			colorRgb := m.At(j, i-1)
			if isDiffColor(bgColor, colorRgb) {
				chMaxY <- i - 1 + line
				return
			}
		}
	}
	chMaxY <- dy - 1
	return
}

func topMinY(m image.Image, bgColor color.Color, dx, dy int, chMinY chan int) {
	for i := 0; i < dy; i = i + line {
		for j := 0; j < dx; j++ {
			colorRgb := m.At(j, i)
			if isDiffColor(bgColor, colorRgb) {
				chMinY <- i + 1 - line
				return
			}
		}
	}
	chMinY <- 0
	return
}

//比较颜色是否相同
func isDiffColor(rgb0, rgb color.Color) bool {
	r0, g0, b0, _ := rgb0.RGBA()
	r, g, b, _ := rgb.RGBA()

	r_uint8 := uint8(r >> 8)
	g_uint8 := uint8(g >> 8)
	b_uint8 := uint8(b >> 8)

	r0_uint8 := uint8(r0 >> 8)
	g0_uint8 := uint8(g0 >> 8)
	b0_uint8 := uint8(b0 >> 8)

	n := param
	if intAbs(r_uint8, r0_uint8) >= n || intAbs(g_uint8, g0_uint8) >= n || intAbs(b_uint8, b0_uint8) >= n {
		return true
	}
	return false
}

func intAbs(a, b uint8) int {
	c := int(a) - int(b)

	if c < 0 {
		return 0 - c
	}
	return c
}
