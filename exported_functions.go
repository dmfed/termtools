// Package termtools
package termtools

import (
	"fmt"
)

// GetColorCodeByID accepts int (from 0 to 255) and returns ANSI escape sequence for requested color. If id is out of range, returns an empty string.
func GetColorCodeByID(id int) (string, error) {
	return getColorByID(id)
}

// GetBackgroundCodeByID accepts int (from 0 to 255) and returns ANSI escape sequence for requested background color. If id is out of range, returns an empty string.
func GetBackgroundCodeByID(id int) (string, error) {
	return getBackgroundByID(id)
}

// GetColorCodeByName accepts color name (string) and returns ANSI escape sequence for requested color. If name is invalid, returns an empty string and error.
func GetColorCodeByName(colorname string) (string, error) {
	return getColorByName(colorname)
}

// GetBackgroundCodeByName accepts color name (string) and returns ANSI escape sequence for requested backgrount color. If name is invalid, returns an empty string and error.
func GetBackgroundCodeByName(colorname string) (string, error) {
	return getBackgroundByName(colorname)
}

// ColorSprint formats using the default formats for its operands and returns the resulting string. If accepts colorname (basic 16 colors).
func ColorSprint(colorname string, a ...interface{}) string {
	return colorSprint(colorname, a...)
}

// ColorIDSprint formats using the default formats for its operands and returns the resulting string. If accepts color id in range [0;255].
func ColorIDSprint(colorcode int, a ...interface{}) string {
	return colorIDSprint(colorcode, a...)
}

// ClearScreen clears screen
func ClearScreen() {
	fmt.Print(Clear)
}

// ClearScreenUp clears screen from current cursor position up
func ClearScreenUp() {
	fmt.Print(ClearUp)
}

// ClearScreenDown clears screen from current cursor position down
func ClearScreenDown() {
	fmt.Print(ClearDown)
}

// ClearLine deletes the whole line of text
func ClearLine() {
	fmt.Print(ClearL)
}

// ClearLineLeft deletes line left of cursor position
func ClearLineLeft() {
	fmt.Print(ClearLLeft)
}

// ClearLineRight deletes line right of cursor position
func ClearLineRight() {
	fmt.Print(ClearLRight)
}

// GetTermSize rturns Current terminal size (in characters)
func GetTermSize() (int, int) {
	return getTermSize()
}

// MoveCursor moves cursor to the specified position on terminal screen. Will do nothing if x or y are out of bounds or we could not get size of terminal
func MoveCursor(x, y int) {
	moveCursor(x, y)
}

// PrintAtPositionAndReturn moves cursor in current terminal to the specified position, prints the string then returns to the original position (when the function was called). Will print at current cursor position if terminal size is unavailable
func PrintAtPositionAndReturn(x, y int, a ...interface{}) {
	printAtPositionAndReturn(x, y, a...)
}
