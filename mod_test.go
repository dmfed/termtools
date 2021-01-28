package termtools

import (
	"fmt"
	"testing"
)

func PrintPalette(width int) {
	p := Printer{}
	colorid := 0
	for colorid < 256 {
		for x := 0; x < width; x++ {
			p.Printf("%3d ", colorid)
			colorid++
			if colorid == 256 {
				colorid += width - x - 1
				break
			}
		}
		colorid -= width
		p.Println()
		for x := 0; x < width; x++ {
			p.SetBackgroundID(colorid)
			p.Printf("   ")
			p.Reset()
			p.Print(" ")
			colorid++
			if colorid == 256 {
				p.Reset()
				break
			}
		}
		p.Println()
	}
}

func Test_basicColors(t *testing.T) {
	var p Printer
	p.Println("Testing basic colors output")
	width := 8
	for color := range colorMap {
		if width == 0 {
			p.Println()
			width = 8
		}
		p.SetColor(color)
		p.Print(color)
		p.Reset()
		width--
	}
	p.Print("\n\n")
}

func Test_PrintPalette(t *testing.T) {
	fmt.Println("These are color codes to set up colors by id...")
	PrintPalette(20)
}
