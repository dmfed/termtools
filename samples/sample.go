package main

import (
	"fmt"
	"math/rand"
	"time"

	tt "github.com/dmfed/termtools"
)

// PrintStair is an example function
func PrintStair(a ...interface{}) {
	rand.Seed(time.Now().UnixNano())
	var p tt.Printer            // initializes Printer
	_, y, _ := tt.GetTermSize() // gets size of terminal window
	y = y/2 - 6
	for xpos := 0; xpos <= 80; xpos = xpos + 5 {
		p.SetColorID(rand.Intn(256))
		tt.PrintAtPosition(xpos, y, p.Sprint(a))
		y++
	}
	p.Reset()
}

func main() {
	// Simplest use is to call ColorSprint (same signature as in fmt.Sprint, but the
	// first argument is the name of the color).
	tt.ClearScreen()                                            // Clears Screen
	tt.MoveCursorTo(0, 0)                                       // Moves cursor to top left
	mystring := tt.ColorSprint("red", "This will print in red") // Colorizes input string
	fmt.Println(mystring)

	// Now using the Printer
	var p tt.Printer                                        // Initialize new Printer
	p.ToggleUnderline()                                     // Set Printer to print underlined
	p.SetColor("green")                                     // Set font color to green
	p.Println("Printing in green underlined with Printer.") // This will print in green underlined

	PrintStair("I'm a walrus!")
}
