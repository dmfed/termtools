package main

import (
	"flag"

	"github.com/dmfed/termtools"
)

// PrintPalette outputs 256 colors with their IDs
func PrintPalette(width int) {
	x, _, _ := termtools.GetTermSize()
	if width < 1 || (width*3+width) > x { // default to 80 columns if in doubt
		width = 20
	}
	p := termtools.Printer{}
	start := 0
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
	p.Println()
}

func main() {
	var width = flag.Int("w", 20, "How many color samples to print on a line?")
	flag.Parse()
	PrintPalette(*width)
}
