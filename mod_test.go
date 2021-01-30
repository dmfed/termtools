package termtools

import (
	"testing"
)

func PrintPalette(width int) {

}

func Test_basicColors(t *testing.T) {
	var p Printer
	p.Println("Testing basic colors output")
	width := 4
	for color := range colorMap {
		if width == 0 {
			p.Println()
			width = 4
		}
		p.SetColor(color)
		p.Printf("[%v] ", color)
		p.Reset()
		width--
	}
	p.Print("\n\n")
}

func Test_ColorOutput(t *testing.T) {
	p := Printer{}
	p.Println("These are color codes to set up colors by id...")
	start, width := 0, 20
	for start < 256 {
		end := start + width
		if end > 256 {
			end = 256
		}
		for id := start; id < end; id++ {
			p.Printf("%3d ", id)
		}
		p.Println()
		for id := start; id < end; id++ {
			p.SetBackgroundID(id)
			p.Printf("   ")
			p.Reset()
			p.Print(" ")
		}
		start = end
		p.Println()
	}
}
