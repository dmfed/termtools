// Package termtools provides basic functionality to style console output on Linux
//
// Copyright 2021 Dmitry Fedotov
package termtools

import (
	"fmt"
)

// GetColorCodeByID accepts int (from 0 to 255 inclusive) and returns ANSI escape
// sequence for requested color. If id is out of range, returns an empty string.
func GetColorCodeByID(colorID int) (string, error) {
	return getColorByID(colorID)
}

// GetBackgroundCodeByID accepts int (from 0 to 255 inclusive) and returns ANSI escape
// sequence for requested background color. If id is out of range, returns an empty string.
func GetBackgroundCodeByID(colorID int) (string, error) {
	return getBackgroundByID(colorID)
}

// GetColorCodeByName accepts color name (string) and returns ANSI escape sequence
// for requested color. If name is invalid, returns an empty string and error.
func GetColorCodeByName(colorname string) (string, error) {
	return getColorByName(colorname)
}

// GetBackgroundCodeByName accepts color name (string) and returns ANSI escape sequence
// for requested backgrount color. If name is invalid, returns an empty string and error.
func GetBackgroundCodeByName(colorname string) (string, error) {
	return getBackgroundByName(colorname)
}

// ColorSprint formats using the default formats for its operands and returns the resulting string.
// It accepts colorname (basic 16 colors). If colorname is invalid will return default format.
func ColorSprint(colorname string, a ...interface{}) string {
	return colorSprint(colorname, a...)
}

// ColorSprintf formats according to a format specifier and returns the resulting string.
// It accepts colorname (basic 16 colors).
func ColorSprintf(colorname string, format string, a ...interface{}) string {
	return colorSprintf(colorname, format, a...)
}

// ColorIDSprint formats using the default formats for its operands and returns the resulting string.
// It accepts color id in range [0;255]. If id is invalid will return default format.
func ColorIDSprint(colorID int, a ...interface{}) string {
	return colorIDSprint(colorID, a...)
}

// ColorIDSprintf formats according to a format specifier and returns the resulting string.
// It accepts color id in range [0;255].
func ColorIDSprintf(colorID int, format string, a ...interface{}) string {
	return colorIDSprintf(colorID, format, a...)
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

// GetTermSize returns Current terminal size (in characters)
// it may fail to get correct values and will return -1, -1 in
// this case. If you're relying on output always check error.
func GetTermSize() (int, int, error) {
	return getTermSize()
}

// MoveCursorTo moves cursor to the specified position on terminal screen.
// Will do nothing if x or y are out of bounds or we can not get size of terminal.
func MoveCursorTo(x, y int) {
	moveCursorTo(x, y)
}

// MoveCursorUp moves cursor up specified number of rows
func MoveCursorUp(rows int) {
	moveCursorUp(rows)
}

// MoveCursorDown moves cursor down specified number of rows
func MoveCursorDown(rows int) {
	moveCursorDown(rows)
}

// MoveCursorLeft moves cursor left specified number of columns
func MoveCursorLeft(columns int) {
	moveCursorLeft(columns)
}

// MoveCursorRight moves cursor left specified number of columns
func MoveCursorRight(columns int) {
	moveCursorRight(columns)
}

// MoveCursorToNextRow moves cursor to next row
func MoveCursorToNextRow() {
	moveCursorToNextRow()
}

// MoveCursorToRow places cursor at the beginning of specified row
func MoveCursorToRow(row int) {
	moveCursorToRow(row)
}

// SaveCursorPosition saves current cursor position.
// Call RestoreCursorPosition() to return
func SaveCursorPosition() {
	saveCursorPosition()
}

// RestoreCursorPosition places cursor to original position when
// SaveCursorPosition was called
func RestoreCursorPosition() {
	restoreCursorPosition()
}

// PrintAtPositionAndReturn moves cursor in the current terminal to the specified position, prints and
// then returns cursor to the inital position (when the function was called).
// Will print at current cursor position if terminal size is unavailable or supplied  x and y
// are out of range (x > terrminal columns or y > terminal rows).
func PrintAtPositionAndReturn(column, row int, a ...interface{}) {
	printAtPositionAndReturn(column, row, a...)
}

// PrintAtPosition moves cursor in the current terminal to the specified position and prints.
// It does not return the cursor to the initial position so subsequent call to
// Print/Println etc. will output immediately after the printed out characters.
// Will print at current cursor position if terminal size is unavailable or supplied  x and y
// are out of range (x > terminal columns or y > terminal rows).
func PrintAtPosition(column, row int, a ...interface{}) {
	printAtPosition(column, row, a...)
}
