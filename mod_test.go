package termtools

import (
	"fmt"
	"testing"
)

func Test_Printer(t *testing.T) {
	p := Printer{}
	p.Println("Testing available 256 colors.")
	for i := 0; i < 256; i++ {
		p.SetColorID(i)
		p.Print("n")
	}
	p.Println()
	p.Reset()
	p.Println("Testing 8 basic colors + 8 bright colors")
	for color := range colorMap {
		p.Reset()
		p.Print(color, " ")
		p.SetColor(color)
		p.Println("nn")
	}
	p.Println()
	p.Reset()
	p.Println("Testing 256 available backrounds")
	for i := 0; i < 256; i++ {
		p.SetBackgroundID(i)
		p.Print(" ")
	}
	p.Println()
	p.Reset()
	p.Println("Testing 8 basic backgrounds + 8 bright backgrounds")
	for color := range colorMap {
		p.SetBackground(color)
		p.Print(" ")
	}
	p.Println()
	p.Reset()
}

func PrintPalette() {
	p := Printer{}
	colorid := 0
	width := 20
	for colorid < 256 {
		for x := 0; x < width; x++ {
			p.Printf("%3d ", colorid)
			colorid++
		}
		colorid -= width
		p.Println()
		for x := 0; x < width; x++ {
			p.SetBackgroundID(colorid)
			p.Printf("   ")
			p.Reset()
			p.Print(" ")
			colorid++
		}
		p.Println()
		p.Reset()
	}
}

func Test_PrintPalette(t *testing.T) {
	fmt.Println("Testing PrintPalette() function")
	PrintPalette()
}

func Test_Orange(t *testing.T) {
	p := Printer{}
	p.SetColorID(7)
	p.SetBackgroundID(9)
	p.Println("This white on orange background")
	p.ToggleReversed()
	p.Println("This is reversed white on orange background")
}
