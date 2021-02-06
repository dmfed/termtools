// Package termtools provides basic functionality to style console output on Linux
//
// v1.0.0
//
// Copyright 2021 Dmitry Fedotov
//
// See package README for usage examples: github.com/dmfed/termtools
//
// General note concerning module usage:
// Whenever color value of type interface{} is required by signature
// of a method or function, either string or int may be supplied.
// Valid color names (to be passed as string) are: "black", "blue", "cyan", "green", "magenta",
// "red", "white", "yellow", "brightblack", "brightblue", "brightcyan",
// "brightgreen", "brightmagenta", "brightred", "brightwhite", "brightyellow".
// Valid color IDs (to be passed as int) are from 0 to 255 inclusive.
//
package termtools

import (
	"fmt"
)

// GetColorCode accepts color identifier (string or int) and returns ANSI escape sequence
// for requested color. If color is invalid the function will return
// empty string and an error.
func GetColorCode(color interface{}) (string, error) {
	return getColorCode(color)
}

// GetBackgroundCode accepts color identifier (string or int) and returns ANSI escape sequence
// for requested background color. If color is invalid the function will return
// empty string and an error.
func GetBackgroundCode(color interface{}) (string, error) {
	return getBackgroundCode(color)
}

// Csprint formats using the default formats for its operands and returns the resulting string.
// It accepts color color identifier (string or int). If color is invalid the function will return
// fmt.Sprint(a).
func Csprint(color interface{}, a ...interface{}) string {
	return colorSprint(color, a...)
}

// Csprintf formats according to a format specifier and returns the resulting string.
// It accepts color color identifier (string or int). If color is invalid the function will return
// fmt.Sprintf(format, a).
func Csprintf(color interface{}, format string, a ...interface{}) string {
	return colorSprintf(color, format, a...)
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

// GetTermSize returns current terminal size (number of columns and rows).
// It may fail to get correct values and will return -1, -1 in
// this case. If you're relying on output to precisely position cursor on screen
// always check error.
func GetTermSize() (columns int, rows int, err error) {
	return getTermSize()
}

// MoveCursorTo moves cursor to the specified position in terminal. (0, 0) is upper left.
// Will do nothing if x or y are out of bounds or we can not get size of terminal.
func MoveCursorTo(column, row int) {
	moveCursorTo(column, row)
}

// MoveCursorHome moves cursor to the upper left corner of the screen.
// Essentially the same as MoveCursorTo(0, 0).
func MoveCursorHome() {
	moveCursorHome()
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

// SaveCursorPosition saves current cursor position and attributes.
// Call RestoreCursorPosition() to return
func SaveCursorPosition() {
	saveCursorPosition()
}

// RestoreCursorPosition places cursor to original position when
// SaveCursorPosition was called and restores attributes.
func RestoreCursorPosition() {
	restoreCursorPosition()
}

// PrintAtPositionAndReturn moves cursor in the current terminal to the specified position, prints, and
// then returns cursor to the inital position.
// Will print at current cursor position if terminal size is unavailable or supplied column and row
// are out of range.
func PrintAtPositionAndReturn(column, row int, a ...interface{}) {
	printAtPositionAndReturn(column, row, a...)
}

// PrintAtPosition moves cursor in the current terminal to the specified position and prints.
// It does not return the cursor to the initial position so subsequent call to
// Print/Println etc. will output immediately after the previous output.
// Will print at current cursor position if terminal size is unavailable or supplied column and row
// are out of range.
func PrintAtPosition(column, row int, a ...interface{}) {
	printAtPosition(column, row, a...)
}
