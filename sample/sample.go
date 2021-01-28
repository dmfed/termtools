package main

import (
	"fmt"

	"github.com/dmfed/termtools"
)

func main() {
	// How to use Printer
	var p termtools.Printer         // initializes Printer
	p.SetColorID(9)                 // Sets color to orange
	termtools.ClearScreen()         //clears screen
	x, y := termtools.GetTermSize() //gets size of terminal window
	s := p.Sprint("We are gophers")
	// Below line outputs string s in the middle of screen and returns cursor to initial position
	termtools.PrintAtPositionAndReturn(x/2-len(s)/2, y/2, s)
	p.ToggleUnderline()
	p.SetColor("green")
	p.Println("We love Go")

	// Simplest use is to call ColorSprint (same signature as in fmt.Sprint)
	mystring := termtools.ColorSprint("red", "This will print in red")
	fmt.Println(mystring)
}
