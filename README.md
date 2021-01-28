## module termtools

**termtools** is basically a collection utilities to style terminal
output and also some utility functions to move cursor around, clear screen, delete lines of text etc.

**go get -u github.com/dmfed/termtools** to use in your code.

First of all you can use can use ANSI escapes directly: attach prefix of color code and suffix of color reset to your string like this:

*s := termtools.Red + "We're gophers" + termtools.Reset*

More common and clever way would be to use Printer

```go
var printer Printer
printer.SetColorID(9)
printer.Println("This will output in orange")
```

Two methods to set color are by name of color and by color numeric id ranging from 0 to 255 inclusive. 

Color numeric ids (color codes) are as follows:

![image of palette with colors numbered 0-255](https://github.com/dmfed/termtools/blob/master/palette.png)

You can see the above palette in your terminal. Just cd into the repo and issu issue **go test**. 

Supported color names are:

**black, red, green, yellow, blue, magenta, cyan, white, brightblack, brightred, brightgreen, brightyellow, brightblue, brightmagenta, brightcyan, brightwhite**

*Note that some of these colors are know to not display correctly in some shells and terminals.*

Printer type implements most print methods in the same way as fmt package from standard library. 
You can call Println, Print, Sprint, Sprintf, Errorf etc. Methods implemented by Printer have same signatures as those if fmt module. In fact termtools just wraps around fmt methods only adding required ANSI escapes to the original values passed and does nothing more than that. 

For a detailed listing of functions and methods [see package documention at pkg.go.dev](https://pkg.go.dev/github.com/dmfed/termtools@v0.6.0)

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
	p.Println("We love Go") 	// This will print in underlined green font

	// Simplest use is to call ColorSprint (same signature as in fmt.Sprint)
	mystring := termtools.ColorSprint("red", "This will print in red")
	fmt.Println(mystring)
}
```
