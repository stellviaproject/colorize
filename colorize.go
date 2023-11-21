package colorize

import (
	"image"
	"image/color"
	"sync"
)

func Colorize(start, end, step float64, list []color.Color, mfs [][]float64) image.Image {
	w, h := len(mfs[0]), len(mfs)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	end += step
	index := 0
	wg := sync.WaitGroup{}
	p := make(chan int, 8)
	for start < end && index < len(list) {
		wg.Add(1)
		go func(min, max float64, index int) {
			defer wg.Done()
			p <- 0
			DrawColorF(min, max, list[index], img, mfs)
			<-p
		}(start, start+step, index)
		start += step
		index++
	}
	wg.Wait()
	return img
}

func DrawColorF(min, max float64, color color.Color, img image.Image, mfs [][]float64) {
	rgba := img.(*image.RGBA)
	w, h := len(mfs[0]), len(mfs)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			v := mfs[y][x]
			if min <= v && v < max {
				rgba.Set(x, y, color)
			}
		}
	}
}

func GenColorList(count int) []color.Color {
	const (
		R_COLOR = 0xFF_00_00
		G_COLOR = 0x00_FF_00
		B_COLOR = 0x00_00_FF
		RANGE   = R_COLOR - B_COLOR
	)
	offset := RANGE / count
	list := make([]color.Color, 0, count)
	for c := B_COLOR; c <= R_COLOR; c += offset {
		r, g, b := c>>16, (c&G_COLOR)>>8, c&B_COLOR
		list = append(list, color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: 255,
		})
	}
	return list
}

var ColorList = []color.Color{
	Yellow,   //1
	Green,    //2
	Blue,     //3
	Red,      //4
	Purple,   //5
	Brown,    //6
	Orange,   //7
	Pink,     //8
	Cyan,     //9
	DarkGray, //10
	LightGray,
	Gray,
	White,
	Olive,
	Rust,
	Salmon,
	Mustard,
	Terracotta,
	Violet,
	Teal,
	Beige,
	Aquamarine,
	Cream,
	Coral,
	Turquoise,
	Gold,
	Silver,
	Black,
	Lavender,
	Lime,
	Burgundy,
}

var (
	LightGray  color.Color = color.RGBA{211, 211, 211, 255}
	DarkGray   color.Color = color.RGBA{169, 169, 169, 255}
	Yellow     color.Color = color.RGBA{255, 255, 0, 255}
	Blue       color.Color = color.RGBA{0, 0, 255, 255}
	White      color.Color = color.RGBA{255, 255, 255, 255}
	Black      color.Color = color.RGBA{0, 0, 0, 255}
	Red        color.Color = color.RGBA{255, 0, 0, 255}
	Green      color.Color = color.RGBA{0, 128, 0, 255}
	Orange     color.Color = color.RGBA{255, 165, 0, 255}
	Purple     color.Color = color.RGBA{128, 0, 128, 255}
	Pink       color.Color = color.RGBA{255, 192, 203, 255}
	Gray       color.Color = color.RGBA{128, 128, 128, 255}
	Brown      color.Color = color.RGBA{165, 42, 42, 255}
	Turquoise  color.Color = color.RGBA{64, 224, 208, 255}
	Gold       color.Color = color.RGBA{255, 215, 0, 255}
	Silver     color.Color = color.RGBA{192, 192, 192, 255}
	Beige      color.Color = color.RGBA{245, 245, 220, 255}
	Aquamarine color.Color = color.RGBA{127, 255, 212, 255}
	Burgundy   color.Color = color.RGBA{128, 0, 32, 255}
	Cyan       color.Color = color.RGBA{0, 255, 255, 255}
	Coral      color.Color = color.RGBA{255, 127, 80, 255}
	Cream      color.Color = color.RGBA{255, 253, 208, 255}
	Lavender   color.Color = color.RGBA{230, 230, 250, 255}
	Lime       color.Color = color.RGBA{0, 255, 0, 255}
	Mustard    color.Color = color.RGBA{255, 219, 88, 255}
	Olive      color.Color = color.RGBA{128, 128, 0, 255}
	Rust       color.Color = color.RGBA{183, 65, 14, 255}
	Salmon     color.Color = color.RGBA{250, 128, 114, 255}
	Teal       color.Color = color.RGBA{0, 128, 128, 255}
	Terracotta color.Color = color.RGBA{204, 78, 92, 255}
	Violet     color.Color = color.RGBA{238, 130, 238, 255}
)

var ColorNames = []string{
	"Yellow",
	"Green",
	"Blue",
	"Red",
	"Purple",
	"Brown",
	"Orange",
	"Pink",
	"Cyan",
	"DarkGray",
	"LightGray",
	"Gray",
	"White",
	"Olive",
	"Rust",
	"Salmon",
	"Mustard",
	"Terracotta",
	"Violet",
	"Teal",
	"Beige",
	"Aquamarine",
	"Cream",
	"Coral",
	"Turquoise",
	"Gold",
	"Silver",
	"Black",
	"Lavender",
	"Lime",
	"Burgundy",
}
