# module termtools

**termtools** is basically a collection of utilities to style console output.
The module allows to colorize terminal output in a very simple and straighforward way.
Basically all of the code is a wrapper around fmt module from Go's standard library. 

**[See full package documention at pkg.go.dev](https://pkg.go.dev/github.com/dmfed/termtools)**

## Download termtools

**go get -u github.com/dmfed/termtools** to use in your code.

## Usage

First of all you can use can use ANSI escapes directly: attach prefix of color escape and suffix of color reset escape  to your string like this:

*s := termtools.Red + "We're gophers" + termtools.Reset*

More common and clever way would be to use Printer

```go
var printer Printer 							// Initializes new Printer
printer.SetColorID(9) 							// Sets color to orange (see color codes below)
printer.Println("This will output in orange") 	// Prints input string in orange
```

Two methods to set color of printer font are by name of color and by color numeric id ranging from 0 to 255 inclusive. 

Color numeric ids (color codes) are as follows:

![image of palette with colors numbered 0-255](https://github.com/dmfed/termtools/blob/main/palette.png)

You can see the above palette in your terminal. Just cd into the repo and issue **go test** It will output 
the pallette with codes. Or issue **go run sample/palette.go** 

Supported color names are:

**black, red, green, yellow, blue, magenta, cyan, white, brightblack, brightred, brightgreen, brightyellow, brightblue, brightmagenta, brightcyan, brightwhite**

*Note that some of these are known to not display correctly in some shells and terminals depending on your settings.*

Printer type implements most print methods in the same way as fmt package from standard library. 
You can call Println, Print, Sprint, Sprintf, Errorf etc. Methods implemented by Printer have same signatures as those if fmt module. In fact termtools just wraps around fmt methods only adding required ANSI escapes to the original values passed and does nothing more than that.

For example fhe following line adds color escape codes to the input:
```go
p := termtools.Printer{} 		// Initializes now Printer
p.SetColor("green") 			// Sets Printer color to green
s := p.Sprint("Hello, world!") 	// s now holds "Hello, world!" with color prefix and reset suffix attached.
```
Note that **len(p.Sprint("Hello, world!"))** will not be the same as **len(fmt.Sprint("Hello, world!"))** because ANSI escapes actually add to the length of the output string. This might be annoying if you're trying
to keep output centered horizontally.


For a detailed listing of module functions and types **[see package documention at pkg.go.dev](https://pkg.go.dev/github.com/dmfed/termtools)**

A demo program is included in **sample** directory of the repo. Here's the listing:

```go
package main

import (
	"fmt"
	"strings"

	tt "github.com/dmfed/termtools"
)

// PrintCentered is an example function
func PrintCentered(a ...interface{}) {
	var p tt.Printer            // initializes Printer
	p.SetColorID(9)             // Sets color to orange
	_, y, _ := tt.GetTermSize() // gets size of terminal window

	// The following line adds color escape codes to the input.
	// Note that the len(p.Sprint(a...)) will not be the same as
	// len(fmt.Sprint(a...)) because ANSI escapes actually add to the
	// length of the output string. This might be annoying if you're trying
	// to keep output centered horizontally.
	s := p.Sprint(a...)

	// below line prints s in the middle of screen and returns cursor to initial position
	tt.PrintAtPosition(0, y/2, s)
	p.Print(strings.Repeat("\n", y/2))
}

func main() {
	// Simplest use of termtools is to call ColorSprint (same signature as in fmt.Sprint, but the
	// first argument is the name of the color).
	tt.ClearScreen()                                            // Clears Screen
	tt.MoveCursor(0, 0)                                         // Moves cursor to top left
	mystring := tt.ColorSprint("red", "This will print in red") // Colorizes input string
	fmt.Println(mystring)

	// Now using the Printer
	var p tt.Printer                                        // Initialize new Printer
	p.ToggleUnderline()                                     // Set Printer to print underlined
	p.SetColor("green")                                     // Set font color to green
	p.Println("Printing in green underlined with Printer.") // This will print in green underlined

	PrintCentered("Printing in the middle of the screen.")
}

```
