package gocaptcha

import (
	"image/color"
	"math"
	"math/rand"
)

//随机生成深色系.
func randDeepColor() color.RGBA {

	randColor := randColor()

	increase := float64(30 + rand.Intn(255))

	red := math.Abs(math.Min(float64(randColor.R)-increase, 255))

	green := math.Abs(math.Min(float64(randColor.G)-increase, 255))
	blue := math.Abs(math.Min(float64(randColor.B)-increase, 255))

	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//随机生成浅色.
func RandLightColor() color.RGBA {

	red := rand.Intn(55) + 200
	green := rand.Intn(55) + 200
	blue := rand.Intn(55) + 200

	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//生成随机颜色.
func randColor() color.RGBA {

	red := rand.Intn(255)
	green := rand.Intn(255)
	blue := rand.Intn(255)
	if (red + green) > 400 {
		blue = 0
	} else {
		blue = 400 - green - red
	}
	if blue > 255 {
		blue = 255
	}
	return color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: uint8(255)}
}

//生成随机字体.
func RandText(num int) string {
	textNum := len(txtChars)
	text := ""
	for i := 0; i < num; i++ {
		text = text + string(txtChars[rand.Intn(textNum)])
	}
	return text
}

// 颜色代码转换为RGB
//input int
//output int red, green, blue.
func ColorToRGB(colorVal int) color.RGBA {

	red := colorVal >> 16
	green := (colorVal & 0x00FF00) >> 8
	blue := colorVal & 0x0000FF

	return color.RGBA{
		R: uint8(red),
		G: uint8(green),
		B: uint8(blue),
		A: uint8(255),
	}
}
