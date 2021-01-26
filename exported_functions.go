package termtools

import (
	"fmt"
	"strings"
)

// GetColorByID accepts int (from 0 to 255) and returns ANSI escape sequence
// for requested color. If id is out of range, returns an empty string.
func GetColorByID(id int) (string, error) {
	return getColorByID(id)
}

// GetBackgroundByID accepts int (from 0 to 255) and returns ANSI escape sequence
// for requested background color. If id is out of range, returns an empty string.
func GetBackgroundByID(id int) (string, error) {
	return getBackgroundByID(id)
}

// GetColorByName accepts color name (string) and returns ANSI escape sequence
// for requested color. If name is invalid, returns an empty string and error.
func GetColorByName(colorname string) (string, error) {
	return getColorByName(colorname)
}

// GetBackgroundByName accepts color name (string) and returns ANSI escape sequence
// for requested backgrount color. If name is invalid, returns an empty string and error.
func GetBackgroundByName(colorname string) (string, error) {
	return getBackgroundByName(colorname)
}

// Paint accepts color name (basic 16 colors) and input string. It then
// appends color prefix and color reset suffix to the string.
func Paint(colorname, input string) (prefix string) {
	colorname = strings.ToLower(colorname)
	prefix, err := getColorByName(colorname)
	if err != nil {
		return input
	}
	return prefix + input + Reset
}

// PaintByID applies color with id colorcode to string
func PaintByID(colorcode int, input string) (prefix string) {
	prefix, err := getColorByID(colorcode)
	if err != nil {
		return input
	}
	return prefix + input + Reset
}

// BPaintByID applies background color with id colorcode to string
func BPaintByID(colorcode int, input string) string {
	prefix, err := getBackgroundByID(colorcode)
	if err != nil {
		return input
	}
	return prefix + input + Reset
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

// ClearLineRight deletes line rigth of cursor position
func ClearLineRight() {
	fmt.Print(ClearLRight)
}

// GetTermSize rturns Current terminal size (in characters)
func GetTermSize() (int, int) {
	return getTermSize()
}

// MoveCursor moves cursor to the specified position on terminal screen
// Will do nothing if x or y are out of bounds or we could not get size of
// terminal
func MoveCursor(x, y int) {
	moveCursor(x, y)
}

// PrintAtPositionAndReturn moves cursor in current terminal to
// the specified position, prints the string then returns to
// the original position (when the function was called)
// Will print at current cursor position if terminal size is unavailable
func PrintAtPositionAndReturn(x, y int, s string) {
	printAtPositionAndReturn(x, y, s)
}
