## module termtools

**termtools** is basically a collection utilities to style terminal
output and also some utility functions to move cursor around, clear screen, delete lines of text etc.

First of all you can use can use ANSI escapes directly: attach prefix of color code and suffix of color reset to your string like this:

*s := termtools.Red + "We're gophers" + termtools.ColorReset*

More common and clever way would be to use Printer

```go
var pinter Printer
printer.SetColorID(9)
printer.Println("This will output in orange")
```

Printer implements most print methods in the same way as fmt package from standard library. 
You can call Println, Print, Sprint, Sprintf, Errorf. Methods implemented by Printer have same signatures as those if fmt module. In fact Printer wraps
those methods only adding required ANSI escapes to the original values passed and does nothing more than that. 

For a detailed listing of functions and mthods signatures see:


A demo program is included in **sample** directory. Here's the listing:

```go
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
```
