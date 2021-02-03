package main

import (
	"fmt"
	"strings"

	"github.com/dmfed/termtools"
)

// PrintCentered is a demo function
func PrintCentered(s string) {
	x, y, _ := termtools.GetTermSize()
	termtools.ClearScreen()
	fmt.Print(strings.Repeat("\n", y))
	s = termtools.Csprint("blue", s)
	termtools.PrintAtPositionAndReturn(x/2-len(s)/2, y/2, s)
}

func main() {
	fmt.Println("This will print in the middle of the screen.")
	PrintCentered("This will print in the middle of the screen.")
}
