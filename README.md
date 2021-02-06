# module termtools

version 1.0.0

**termtools** is basically a collection of utilities to style console output on Linux systems.
The module allows to colorize terminal output in a very simple and straighforward way.
Most of the code is a wrapper around fmt module from Go's standard library. 

**termtools** module also includes utility functions to control cursor position in the terminal.

**Note: This README is NOT intended to document all of what termtools has to offer and only provides some usage examples.** 
**Please see full module documentation at** **[https://pkg.go.dev/github.com/dmfed/termtools](https://pkg.go.dev/github.com/dmfed/termtools)** **for a complete list of what is available.**

## Getting termtools

**go get -u github.com/dmfed/termtools** and you are good to Go :) 

**import "github.com/dmfed/termtools"** to import module in you code

## Basic usage

First of all you can directly use any ANSI escapes declared in the module. (For a complete list of exported constants see **[https://pkg.go.dev/github.com/dmfed/termtools](https://pkg.go.dev/github.com/dmfed/termtools)**.)

```go
s := termtools.Red + "We're gophers" + termtools.Reset 	// adding escape for red and escape for color reset
fmt.Println(s) 						// the string "We're gophers" will output in red.
```
You can also get your string colorized using function similar to fmt.Sprint (same signature as Sprint, but first parameter is color name). 

```go
mystring := termtools.Csprint("red", "This will print in red") // Colorizes input string
fmt.Println(mystring)
```
See below for a full list of supported color names. 

## Using Printer

More clever way would be to use Printer. 

```go
var printer termtools.Printer 				// Initializes new Printer
printer.SetColor(9) 					// Sets color (9 is orange)
printer.Println("This will output in orange") 		// This prints in orange
printer.SetColor("red")					// Sets color
printer.Println("This will output in red")		// This prints in red
printer.ToggleUnderline()				// switches undeline mode on
printer.Println("This will output in undelined red")	// prints in underlined red
```

Font and backgrount colors can be set either by color name (supplied as string) or by color numeric id ranging from 0 to 255 inclusive. To set font or background color by name use **printer.SetColor(colorname interface{})** and **printer.SetBackground(colorname string)** methods of Printer accordingly. See **[https://pkg.go.dev/github.com/dmfed/termtools](https://pkg.go.dev/github.com/dmfed/termtools)** for a complete list of Printer methods.

Supported color names are:

**"black", "blue", "cyan", "green", "magenta", "red", "white", "yellow", "brightblack", "brightblue", "brightcyan", "brightgreen", "brightmagenta", "brightred", "brightwhite", "brightyellow"**

Color numeric IDs are as follows:

![image of palette with colors numbered 0-255](https://raw.githubusercontent.com/dmfed/termtools/main/palette.png)

To see the above palette in your terminal **git clone https://github.com/dmfed/termtools/** then **cd** into the repo and issue **go test**. The test will output known colors and the pallette with codes. See also see **samples/palette.go**

*NOTE: colors may not display correctly in some shells and terminals depending on your settings.*

## Printer modes: bold, underline, reversed, blinking
Printer has four modes: **bold**, **reversed**, **underline**, and **blinking**. Bold, underline and blinking are self explanatory. Reversed mode if switched on swaps font and background colors). These modes can be toggled on and off with **ToggleBold()**, **ToggleBlinking()**, **ToggleReversed()**, and **ToggleUnderline()** methods of Printer. For a complete list of Printer methods see **[https://pkg.go.dev/github.com/dmfed/termtools](https://pkg.go.dev/github.com/dmfed/termtools)**.

## NewPrinter
Printer instance can be initialized with all options set to needed values at once. Use **termtools.NewPrinter()** to set up instance of Printer as required.

The functions sugnature is as follows: 

**func NewPrinter(conf PrinterConfig) (p \*Printer, err error)**

It accepts PrinterConfig as single argument.

```go
type PrinterConfig struct {
	Name       string
	Color      interface{}
	Background interface{}
	Bold       bool
	Underline  bool
	Reversed   bool
	Blinking   bool
	Prefix     string
	Suffix     string
}
```

Here is an example.

```go
printer, err := termtools.NewPrinter(termtools.PrinterConfig{Color: "blue", Blinking: true})
if err != nil {
	printer.Println(err)
}
printer.Println("Printing in blue blinking here.")
```
## Printing methods
Printer type implements most print methods as in fmt module from standard library copying their signatures, names and behaviour. You can call **Println, Print, Sprint, Sprintf, Errorf** etc. In fact printing methods just wrap around fmt module by adding required ANSI escapes to the original values passed.

For example fhe following adds color escape codes to the input and color reset escape at the end:
```go
greenprinter := termtools.Printer{} 		// initializes new Printer
greenprinter.SetColor("green") 				// sets Printer color to green
s := greenprinter.Sprint("Hello, world!") 	// s now holds "Hello, world!" with green color prefix and reset suffix attached.
```
Note that **len(greenprinter.Sprint("Hello, world!"))** in the above example will not be the same as **len(fmt.Sprint("Hello, world!"))** because ANSI escapes actually add to the length of the output string. This might be annoying if you're trying to keep output centered horizontally and rely on calculation of string length.

For a detailed list of printing methods **[see package documention at pkg.go.dev](https://pkg.go.dev/github.com/dmfed/termtools)**

## Moving cursor around

**termtools** includes functions and methods of Printer type to manipulate cursor position on screen, clear screen and delete lines of text.

```go
func PrintCentered(s string) {
	x, y, _ := termtools.GetTermSize() 				// returns number of columns and rows in current terminal
	termtools.ClearScreen()						// clears screen
	fmt.Print(strings.Repeat("\n", y))
	s = termtools.ColorSprint("blue", s)				// sets input string color to blue
	termtools.PrintAtPositionAndReturn(x/2-len(s)/2, y/2, s)	// prints at requested position and returns cursor
}
```
Or using **termtools.Printer**
```go
var p termtools.Printer
p.MoveTo(10,10)
p.Print("This starts at column 10, row 10")
```
**termtools** has the following functions to control cursor position (signatures are self-explanatory):

**func MoveCursorTo(column, row int), func MoveCursorHome(), func MoveCursorUp(rows int), func MoveCursorDown(rows int), MoveCursorLeft(columns int), MoveCursorRight(columns int), MoveCursorToRow(row int), func MoveCursorToNextRow()**

Most of this functionality is also implemented for **termtools.Printer** type with shorter and bit different names. **[See docs at pkg.go.dev](https://pkg.go.dev/github.com/dmfed/termtools)**

## Clearing screen, deleting lines of text

**termtools** has the following set of functions available to delete lines of text and clear screen:

**ClearScreen(), ClearScreenUp(), ClearScreenDown(), ClearLine(), ClearLineLeft(), ClearLineRight()**

## Using PrintSuite to style your program output

**termtools.PrintSuite** can act as an (almost) full replacement to fmt module from standard library. It is intended to hold one or more configurations of **termtools.Printer** and switch them on the fly. This way you
can get different output styles for different kinds of messages in your program. 

```go
type PrintSuite struct {
	Printer
	// Has unexported field.
}
```
**termtools.PrintSuite** embeds default Printer, so you can call any printing method on it directly.

In example below we set up PrintSuite to be available globally in the program then configure two printers (red and green underlined). 

We then can use (for example) red printer to display error messages and green printer to display success notifications.

Consider code below:
```go
package main

import "github.com/dmfed/termtools"

var prnt = termtools.PrintSuite{} // here we define prnt globally to reuse it from any part of program

func setupPrinters() {
	configs := []termtools.PrinterConfig{
		{Name: "error", Color: "red"},
		{Name: "notify", Color: "green", Underline: true}}
	if err := prnt.Configure(configs...); err != nil {
		prnt.Println(err)
	}
}

func main() {
	setupPrinters()
	// We now can use different printers for different kinds of messages anywhere in out program.
	// For example print errors in red
	prnt.Use("error").Print("We call prnt.Use(\"error\").Print() to output using printer named \"error\"\n")

	// or print notifications in green
	prnt.Use("notify").Print("and prnt.Use(\"notify\").Print() to use printer named \"notify\"\n")

	// or use unstyled output of embedded printer for everything else like this:
	prnt.Println("This will print without any styling.")

	// or like this:
	prnt.UseDefault().Println("This will print without any styling.")
}
```
**IMPORTANT: When passing termtools.PrinterConfig configuration(s) to PrintSuite Configure() method always make sure that Name field ot config(s) is NOT EMPTY. Otherwise Configure will fail with an error.** Some of passed configurations (precceding the one that is missing name) may be processed and PrintSuite will still be usable but the mathod will terminate upon encountering config with empty Name field.

Note that **Use()** method in example above returns pointer to named printer, so you can call any Printer method directly. 
Also note that PrintSuite embeds Printer instance so any Printer method can be called on it without Use() (i.e **prnt.Println("This will print without any styling.")**). This will use embedded Printer instance.

To change embedded printer style use SwitchTo(printername string).

```go
prnt.SwitchTo("notify")
prnt.Println("Call to prnt.Println (and other methods) from this point on will use \"notify\" style.")
```

To reset embedded Printer instance use **SwitchToDefault()**

```go
prnt.SwitchToDefault()
prnt.Println("Call to prnt.Println will now act as simple call to fmt.Println again."
```
For a full list of PrintSuite methods **[see package documention at pkg.go.dev](https://pkg.go.dev/github.com/dmfed/termtools)**

## Example programs

Some example programs are included in **samples** directory of the repo at https://github.com/dmfed/termtools/tree/main/samples

Here's the listing of one of the examples:

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
Please see **[https://pkg.go.dev/github.com/dmfed/termtools](https://pkg.go.dev/github.com/dmfed/termtools)** for a complete documentation of **termtools**.

## That's all folks!

Thank you for reading this far :)

I hope you might find **termtools** useful. Any feedback is welcome. 

I'm planning to freeze all function signatures in v1.0.0. Until then some functions signatures might change and I'm also considering adding and/or removing some stuff. 
