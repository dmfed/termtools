# module termtools

**termtools** is basically a collection of utilities to style console output on Linux systems.
The module allows to colorize terminal output in a very simple and straighforward way.
Most of the code is a wrapper around fmt module from Go's standard library. 

**termtools** module also includes utility functions to control cursor position in the terminal.

**Note: This README is NOT intended to document all of what termtools has to offer and only provides basic usage examples.** 
**Please see full module documention at** **[https://pkg.go.dev/github.com/dmfed/termtools](https://pkg.go.dev/github.com/dmfed/termtools)** **for a complete list of what is available.**

## Getting termtools

**go get -u github.com/dmfed/termtools** and you are good to Go :) 

## Basic usage

First of all you can directly use any ANSI escapes declared in the module. (For a complete list of exported constants see **[https://pkg.go.dev/github.com/dmfed/termtools](https://pkg.go.dev/github.com/dmfed/termtools)**.)

```go
s := termtools.Red + "We're gophers" + termtools.Reset 	// adding escape for red and escape for color reset
fmt.Println(s) 						// the string "We're gophers" will output in red.
```
You can also get your string colorized using function similar to fmt.Sprint (same signature as Sprint, but first parameter is color name). 

```go
mystring := termtools.ColorSprint("red", "This will print in red") // Colorizes input string
fmt.Println(mystring)
```
See below for a full list of supported color names. 

## Using Printer

More clever way would be to use Printer. 

```go
var printer Printer 					// Initializes new Printer
printer.SetColorID(9) 					// Sets color by id (9 is orange)
printer.Println("This will output in orange") 		// This prints in orange
printer.SetColor("red")					// Sets color by name 
printer.Println("This will output in red")		// This prints in red
printer.ToggleUnderline()				// switches undeline mode on
printer.Println("This will output in undelined red")	// prints in underlined red
```

Font and backgrount colors can be set either by color name (supplied as string) or by color numeric id ranging from 0 to 255 inclusive. To set font or background color by name use **SetColor(colorname string)** and **SetBackground(colorname string)** methods of Printer accordingly. To set color or background by numeric ID use **SetColorID(colorID int)** and **SetBackgroundID(colorID int)** methods of Printer. See **[https://pkg.go.dev/github.com/dmfed/termtools](https://pkg.go.dev/github.com/dmfed/termtools)** for a complete list of Printer methods.

Supported color names are:

**"black", "blue", "cyan", "green", "magenta", "red", "white", "yellow", "brightblack", "brightblue", "brightcyan", "brightgreen", "brightmagenta", "brightred", "brightwhite", "brightyellow"**

Color numeric IDs are as follows:

![image of palette with colors numbered 0-255](https://github.com/dmfed/termtools/blob/main/palette.png)

You can see the above palette in your terminal. Just **git clone https://github.com/dmfed/termtools/** then **cd** into the repo and issue **go test**. The test will output known colors and the pallette with codes. Or see **samples/palette.go**

*NOTE: some of these colors are known to not display correctly in some shells and terminals depending on your settings.*

## Printer modes: bold, underline, reversed
Printer has three modes: **bold**, **reversed** and **underline**. Bold and underline are self explanatory. Reversed mode if switched on makes Printer output font colored with background color and background colored with font color (basically swaps font and background colors). These modes can be toggled on and off with **ToggleBold()**, **ToggleReversed()**, and **ToggleUnderline()** methods of Printer. For a complete list of Printer methods see **[https://pkg.go.dev/github.com/dmfed/termtools](https://pkg.go.dev/github.com/dmfed/termtools)**.

## NewPrinter and NewPrinterID
Printer instance can be initialized with all options set to needed values at once. There are two functions aimed at this:

**func NewPrinter(fontcolor, backgroundcolor string, bold, underline, reversed bool) (\*Printer, error)**

**func NewPrinterID(fontcolorID, backgroundcolorID int, bold, underline, reversed bool) (\*Printer, error)**

```go
printer, err := NewPrinter("red", "", false, true, false) // returns new Printer with font set to red underline
if err != nil {
	// handle error 
}
printer.Print("Hello, world!")
```
## Printing methods
Printer type implements most print methods as in fmt module from standard library mirroring their signatures, names and behaviour. You can call Println, Print, Sprint, Sprintf, Errorf etc. In fact printing methods just wrap around fmt methods by adding required ANSI escapes to the original values passed.

For example fhe following adds color escape codes to the input and color reset escape at the end:
```go
greenprinter := termtools.Printer{} 		// initializes new Printer
greenprinter.SetColor("green") 			// sets Printer color to green
s := greenprinter.Sprint("Hello, world!") 	// s now holds "Hello, world!" with green color prefix and reset suffix attached.
```
Note that **len(p.Sprint("Hello, world!"))** in the above example will not be the same as **len(fmt.Sprint("Hello, world!"))** because ANSI escapes actually add to the length of the output string. This might be annoying if you're trying to keep output centered horizontally and rely on calculation of string length.

For a detailed listing of module functions and types **[see package documention at pkg.go.dev](https://pkg.go.dev/github.com/dmfed/termtools)**

## Moving cursor around

**termtools** includes functions and methods of Printer type to manipulate cursor position on screen, clear screen and delete line of text.

```go
func PrintCentered(s string) {
	x, y, _ := termtools.GetTermSize() 				// returns number of columns and rows in current terminal
	termtools.ClearScreen()						// clears screen
	fmt.Print(strings.Repeat("\n", y))
	s = termtools.ColorSprint("blue", s)				// sets input string color to blue
	termtools.PrintAtPositionAndReturn(x/2-len(s)/2, y/2, s)	// prints at requested position and returns cursor
}
```
Please see **[https://pkg.go.dev/github.com/dmfed/termtools](https://pkg.go.dev/github.com/dmfed/termtools)** for a complete list of functions relating to cursor movement.

## Example program

A couple of example programs are included in **samples** directory of the repo. Here's the listing of one of the examples:

```go
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
		p.PrintAtPosition(xpos, y, a...)
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
```
Thank you for reading this far :)

I hope you might find this module useful. Any feedback is welcome. 
