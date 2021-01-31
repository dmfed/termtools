package termtools

import (
	"fmt"
	"io"
)

// Printer holds color and style settings and implements most methods as in fmt module like Print,
// Println, Sprint etc. adding color and styles to the input values.
type Printer struct {
	color      string
	background string
	bold       bool
	underline  bool
	reversed   bool
}

// NewPrinter takes font and background color names and three bool values (bool, underline, reversed) and
// returns instance of Printer with requested properties set accordingly. Valid color and background names are:
// "black", "blue", "cyan", "green", "magenta", "red", "white", "yellow", "brightblack", "brightblue",
// "brightcyan", "brightgreen", "brightmagenta", "brightred", "brightwhite", "brightyellow".
// If invalid color or background name is supplied then that color or background will not be set and
// the function will return an error ErrUnknownColor. Other requested parameters will be set.
func NewPrinter(fontcolor, backgroundcolor string, bold, underline, reversed bool) (*Printer, error) {
	var p Printer
	var result error
	if code, err := getColorByName(fontcolor); err == nil {
		p.color = code
	} else {
		result = err
	}
	if code, err := getBackgroundByName(backgroundcolor); err == nil {
		p.background = code
	} else {
		result = err
	}
	p.bold = bold
	p.underline = underline
	p.reversed = reversed
	return &p, result
}

// NewPrinterID takes font and background color IDs and three bool values (bool, underline, reversed) and
// returns instance of Printer with requested properties set accordingly. Valid color and background
// IDs are ints in range [0;255]. If invalid color or background ID is supplied then that color or background
// will not be set and the function will return an error ErrUnknownColor. Other requested parameters will
// be set.
func NewPrinterID(fontcolorID, backgroundcolorID int, bold, underline, reversed bool) (*Printer, error) {
	var p Printer
	var result error
	if code, err := getColorByID(fontcolorID); err == nil {
		p.color = code
	} else {
		result = err
	}
	if code, err := getBackgroundByID(backgroundcolorID); err == nil {
		p.background = code
	} else {
		result = err
	}
	p.bold = bold
	p.underline = underline
	p.reversed = reversed
	return &p, result
}

// Errorf formats according to a format specifier and returns the string as a value that satisfies error.
func (p *Printer) Errorf(format string, a ...interface{}) error {
	out := p.processString(format)
	return fmt.Errorf(out, a...)
}

// Fprint formats using the default formats for its operands and writes to w. Spaces are added between
// operands when neither is a string. It returns the number of bytes written and any write error encountered.
func (p *Printer) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	out := p.processString(a...)
	return fmt.Fprint(w, out)
}

// Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written
// and any write error encountered.
func (p *Printer) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	out := p.processString(format)
	return fmt.Fprintf(w, out, a...)
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.
func (p *Printer) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	out := p.processString(a...)
	return fmt.Fprintln(w, out)
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.
func (p *Printer) Print(a ...interface{}) (n int, err error) {
	out := p.processString(a...)
	return fmt.Print(out)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p *Printer) Printf(format string, a ...interface{}) (n int, err error) {
	out := p.processString(format)
	return fmt.Printf(out, a...)
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.
func (p *Printer) Println(a ...interface{}) (n int, err error) {
	out := p.processString(a...)
	return fmt.Println(out)
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p *Printer) Sprint(a ...interface{}) string {
	out := p.processString(a...)
	return fmt.Sprint(out)
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p *Printer) Sprintf(format string, a ...interface{}) string {
	out := p.processString(format)
	return fmt.Sprintf(out, a...)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p *Printer) Sprintln(a ...interface{}) string {
	out := p.processString(a...)
	return fmt.Sprintln(out)
}

// SetColor sets color of printer defined by colorname input string.
// If colorname is not known to the library or is empty the method will return an error.
func (p *Printer) SetColor(colorname string) error {
	code, err := getColorByName(colorname)
	if err == nil {
		p.color = code
	}
	return err
}

// SetColorID sets color of printer defined by id in range [0;255].
// If id is out of range the method will return an error.
func (p *Printer) SetColorID(colorID int) error {
	code, err := getColorByID(colorID)
	if err == nil {
		p.color = code
	}
	return err
}

// SetBackground sets backgtound color of printer defined by colorname input string.
// If colorname is not known or is empty the method will return an error.
func (p *Printer) SetBackground(colorname string) error {
	code, err := getBackgroundByName(colorname)
	if err == nil {
		p.background = code
	}
	return err
}

// SetBackgroundID sets color of printer defined by id in range [0;255].
// If id is out of range the method will return an error.
func (p *Printer) SetBackgroundID(colorID int) error {
	code, err := getBackgroundByID(colorID)
	if err == nil {
		p.background = code
	}
	return err
}

// ToggleBold toggles bold mode of Printer
func (p *Printer) ToggleBold() {
	p.bold = !p.bold
}

// ToggleUnderline toggles underline mode of Printer
func (p *Printer) ToggleUnderline() {
	p.underline = !p.underline
}

// ToggleReversed toggles reverse mode of Printer
func (p *Printer) ToggleReversed() {
	p.reversed = !p.reversed
}

// Reset resets printer state to initial state (no color, no background, bold, underline and reversed modes turned off).
func (p *Printer) Reset() {
	p.color = ""
	p.background = ""
	p.bold = false
	p.underline = false
	p.reversed = false
}

// PrintAtPosition moves cursor to specified column and row and issues Print
// It does not return to the initial position. See also PrintAtPositionAndReturn method.
func (p *Printer) PrintAtPosition(column, row int, a ...interface{}) (n int, err error) {
	out := p.processString(a...)
	moveCursorTo(column, row)
	return fmt.Print(out)
}

// PrintAtPositionAndReturn moves cursor to specified column and row and issues Print
// then moves cursor to initial position when method was called.
func (p *Printer) PrintAtPositionAndReturn(column, row int, a ...interface{}) (n int, err error) {
	out := p.processString(a...)
	saveCursorPosition()
	moveCursorTo(column, row)
	defer restoreCursorPosition()
	return fmt.Print(out)
}

// MoveTo places cursor at the specafied column and row
func (p *Printer) MoveTo(column, row int) {
	moveCursorTo(column, row)
}

// MoveUp moves cursor up specified number of rows
func (p *Printer) MoveUp(rows int) {
	moveCursorUp(rows)
}

// MoveDown moves cursor down specified number of rows
func (p *Printer) MoveDown(rows int) {
	moveCursorDown(rows)
}

// MoveLeft moves cursor left specified number of columns
func (p *Printer) MoveLeft(columns int) {
	moveCursorLeft(columns)
}

// MoveRight moves cursor right specified number of columns
func (p *Printer) MoveRight(columns int) {
	moveCursorRight(columns)
}

// MoveToNextRow moves cursor to the next row
func (p *Printer) MoveToNextRow() {
	moveCursorToNextRow()
}

// MoveToRow moves cursor to the specified row
func (p *Printer) MoveToRow(row int) {
	moveCursorToRow(row)
}

func (p *Printer) processString(a ...interface{}) string {
	out := p.color + p.background
	if p.bold {
		out += Bold
	}
	if p.underline {
		out += Underline
	}
	if p.reversed {
		out += Reversed
	}
	if out == "" {
		return fmt.Sprint(a...)
	}
	out += fmt.Sprint(a...)
	out += Reset
	return out
}
