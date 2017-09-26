package main

import 
(
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	w int;
	h int;
	colr uint8;
	
}

func main() {
	m := Image{500,500,0}
	pic.ShowImage(m)
}


func (a Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (a Image) Bounds() image.Rectangle{
	return image.Rect(0,0,a.w,a.h)
}

func (a Image) At(x, y int) color.Color{
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}


