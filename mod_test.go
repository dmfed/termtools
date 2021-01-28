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

func Test_PrintPalette(t *testing.T) {
	fmt.Println("These are color codes to set up colors by id...")
	PrintPalette(20)
}
