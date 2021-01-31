# module termtools

**termtools** is basically a collection of utilities to style console output.
The module allows to colorize terminal output in a very simple and straighforward way.
Most of the code is a wrapper around fmt module from Go's standard library. 

The module also includes utility functions to control cursor position in terminal window.

**[See full package documention at pkg.go.dev](https://pkg.go.dev/github.com/dmfed/termtools)**

## Download termtools

**go get -u github.com/dmfed/termtools** to use in your code.

## Usage

First of all you can use can use ANSI escapes declared in the module directly like this:

```go
s := termtools.Red + "We're gophers" + termtools.Reset
fmt.Println(s) // the string "We're gophers" will output in red.
```

More common and clever way would be to use Printer:

```go
var printer Printer 									// Initializes new Printer
printer.SetColorID(9) 									// Sets color by id (see below)
printer.Println("This will output in orange") 			// This prints in orange
printer.SetColor("red")									// Sets color by name 
printer.Println("This will output in red")				// This prints in red
printer.ToggleUnderline()								// swithces undeline mode on
printer.Println("This will output in undelined red")	// prints in underlined red
```

Font and backgrount colors can be set either by color name (supplied as string) or by color numeric id ranging from 0 to 255 inclusive. To set font or background color by name use **SetColor(colorname string)** and **SetBackground(colorname string)** methods of Printer accordingly. To set color or background by numeric ID use **SetColorID(colorID int)** and **SetBackgroundID(colorID int)** methods of Printer.

Supported color names are:

**"black", "blue", "cyan", "green", "magenta", "red", "white", "yellow", "brightblack", "brightblue", "brightcyan", "brightgreen", "brightmagenta", "brightred", "brightwhite", "brightyellow"**

Color numeric IDs are as follows:

![image of palette with colors numbered 0-255](https://github.com/dmfed/termtools/blob/main/palette.png)

You can see the above palette in your terminal. Just cd into the repo and issue **go test** It will output 
the pallette with codes. Or issue **go run samples/palette.go**

*Note that some of these are known to not display correctly in some shells and terminals depending on your settings.*

Printer also has three modes: bold, reversed and underline. Bold and underline are self explanatory. Reversed mode if switched on makes Printer output font colored with background color and background colored with font color. These three modes can be toggled on and off with **ToggleBold()**, **ToggleReversed()**, and **ToggleUnderline()** methods of Printer.

Printer instance can be initialized with all options set to needed values at once. There are two functions aimed at this:

**func NewPrinter(fontcolor, backgroundcolor string, bold, underline, reversed bool) (*Printer, error)**

**func NewPrinterID(fontcolorID, backgroundcolorID int, bold, underline, reversed bool) (*Printer, error)**

```go
printer, err := NewPrinter("red", "", false, false, false) // returns new Printer with font set to red
if err != nil {
	// handle error 
}
printer.Print("Hello, world!")
```

Printer type implements most print methods as in fmt module from standard library mirroring their signatures
and names. You can call Println, Print, Sprint, Sprintf, Errorf etc. In fact termtools just wraps around fmt methods only adding required ANSI escapes to the original values passed and does nothing more than that.

For example fhe following line adds color escape codes to the input:
```go
p := termtools.Printer{} 		// Initializes now Printer
p.SetColor("green") 			// Sets Printer color to green
s := p.Sprint("Hello, world!") 	// s now holds "Hello, world!" with color prefix and reset suffix attached.
```
Note that **len(p.Sprint("Hello, world!"))** will not be the same as **len(fmt.Sprint("Hello, world!"))** because ANSI escapes actually add to the length of the output string. This might be annoying if you're trying
to keep output centered horizontally and rely on calculation of string length. 

Printer can also be initialized 

For a detailed listing of module functions and types **[see package documention at pkg.go.dev](https://pkg.go.dev/github.com/dmfed/termtools)**

A demo program is included in **samples** directory of the repo. Here's the listing:

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
