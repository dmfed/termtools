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
