## module termtools

**termtools** is basically a collection of ANSI escape sequences used to style terminal
output and some utility functions to move cursor around, clear screen, delete lines of text etc.


**constants.go** holds ANSI codes of different colors and styles (bold, underline, reversed), color reset code, and codes to save and restore position of cursor.


**functions.go** holds utility functions for simple manipulations.


To colorise your output (either font color or background color) you can attach prefix of color code and suffix of color reset to your string directly like this:


*s := termtools.Red + "We're gophers" + termtools.ColorReset*


or


*s := fmt.Printf("This is %vgreen%v", termtool.Green, termtool.ColorReset)*


Or you can use functions provided like this:


*fmt.Println(termtools.Paint("cyan", "We like Go!"))*


A demo program is included in **sample** directory. Here's the listing:

```
package main

import (
	"fmt"

	"github.com/dmfed/termtools"
)

func main() {
	termtools.ClearScreen()                                            //clears screen
	x, y := termtools.GetCurrTermSize()                                //gets size of terminal window
	s := termtools.Reversed + termtools.Paint("blue", "Hello, world!") //adds blue color to string "Hello,world!"
	termtools.PrintAtPositionAndReturn(x/2-len(s)/2, y/2, s)           //outputs string s in the middle of screen and returns cursor to initial position
	s2 := termtools.Red + "We're gophers" + termtools.ColorReset
	fmt.Println(s2)
	fmt.Println(termtools.Paint("cyan", "We like Go!"))
}
```