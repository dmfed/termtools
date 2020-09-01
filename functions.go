package termtools

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

//ByID accepts int (from 0 to 255) and returns ANSI escape sequence
//suitable for setting color of strings
func ByID(id int) string {
	if id >= 0 && id < 256 {
		return fmt.Sprintf("\u001b[38;5;%vm", id)
	}
	return ""
}

//BByID accepts int (from 0 to 255) and returns ANSI escape sequence
//suitable for setting background color of a string
func BByID(id int) string {
	if id >= 0 && id < 256 {
		return fmt.Sprintf("\u001b[48;5;%vm", id)
	}
	return ""
}

//Reset returns string that resets font background and style
func Reset() string {
	return ColorReset
}

//Paint accepts color name (basic 16 colors) and input string. It then
//appends color prefix and color reset postfix to the string.
func Paint(color, input string) string {
	color = strings.ToLower(color)
	prefix := ""
	switch color {
	case "black":
		prefix = Black
	case "red":
		prefix = Red
	case "green":
		prefix = Green
	case "yellow":
		prefix = Yellow
	case "blue":
		prefix = Blue
	case "magenta":
		prefix = Magenta
	case "cyan":
		prefix = Cyan
	case "white":
		prefix = White
	case "brightblack":
		prefix = BrightBlack
	case "brightred":
		prefix = BrightRed
	case "brightgreen":
		prefix = BrightGreen
	case "brightyellow":
		prefix = BrightYellow
	case "brightblue":
		prefix = BrightBlue
	case "brightmagenta":
		prefix = BrightMagenta
	case "brightcyan":
		prefix = BrightCyan
	case "brightwhite":
		prefix = BrightWhite
	default:
		return input
	}
	return prefix + input + ColorReset
}

//PaintByID applies color returned by ByID to string
func PaintByID(colorcode int, input string) string {
	prefix := ByID(colorcode)
	return prefix + input + ColorReset
}

//BPaintByID applies color returned by ByID to string
func BPaintByID(colorcode int, input string) string {
	prefix := BByID(colorcode)
	return prefix + input + ColorReset
}

//MoveCursor moves cursot to the cpecified position on terminal screen
//Will do nothing if x or y are out of bounds or we could not get size of
//terminal
func MoveCursor(x, y int) {
	maxx, maxy := GetCurrTermSize()
	if x <= maxx && y <= maxy {
		fmt.Printf("\033[%v;%vH", x, y)
	}
}

//PrintAtPositionAndReturn moves cursor in current terminal to
//the specified position, prints the string then returns to
//the original position (when the function was called)
//Will print at current cursor position if terminal size is unavailable
func PrintAtPositionAndReturn(y, x int, s string) {
	fmt.Print(SaveCursor)
	MoveCursor(y, x)
	fmt.Print(s)
	fmt.Print(RestoreCursor)
}

//GetCurrTermSize rturns Current terminal size (in characters)
func GetCurrTermSize() (int, int) {
	x, y, err := terminal.GetSize(0)
	if err == nil {
		return x, y
	}
	return 0, 0
}

//ClearScreen does exactly what its name says
func ClearScreen() {
	fmt.Print("\u001b[2J")
}

func ClearScreenUp() {
	fmt.Print("\u001b[1J")
}

func ClearScreenDown() {
	fmt.Print("\u001b[0J")
}

//ClearLine deletes the whole line of text
func ClearLine() {
	fmt.Print("\u001b[2K")
}

func ClearLineLeft() {
	fmt.Print("\u001b[1K")
}

func ClearLineRight() {
	fmt.Print("\u001b[0K")
}
