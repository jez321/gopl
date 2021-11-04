// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image/color"
	"os"

	"github.com/jez321/gopl/ch01/1.6/lissajous"
)

var palette = []color.Color{color.Black, color.RGBA{0x0, 0xff, 0x0, 0xff}, color.RGBA{0xff, 0xff, 0x0, 0xff}, color.RGBA{0x0, 0xff, 0xff, 0xff}}

func main() {
	lissajous.Write(os.Stdout)
}
