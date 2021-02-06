package termtools

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrFailedToSetColor      = errors.New("error: failed to set printer color")
	ErrFailedToSetBackground = errors.New("error: failed to set printer background")
)

// Printer holds color and style settings and implements most methods as in fmt module like Print,
// Println, Sprint etc. adding color and styles to the input values.
type Printer struct {
	color      string
	background string
	prefix     string
	suffix     string
	bold       bool
	underline  bool
	reversed   bool
	blinking   bool
	name       string
}

// PrinterConfig describes configuration of Printer.
type PrinterConfig struct {
	// Name identifies configuration when used
	// with PrintSuite type. Name field is mandatory when passing PrinterConfig
	// to PrintSuite Configure method!
	//
	// When setting up a Printer instance with NewPrinter() Name field may be omitted.
	Name string
	// Color and Backgrount fields may hold color name (as string) or color ID in range
	// [0;255]. See package docs for list of available color names.
	Color      interface{}
	Background interface{}
	// Bold, Underline, Reversed, Blinking switch relevant printer modes on if set to true.
	Bold      bool
	Underline bool
	Reversed  bool
	Blinking  bool
	// Prefix and suffix are added to output if they are not empty strings.
	Prefix string
	Suffix string
}

// NewPrinter takes PrinterConfig and returns pointer to Printer.
// If invalid color(s) identifier is encountered the function returns
// instance of printer with color(s) not set. Otherwise the printer is functional.
// Check for errors to make sure that returned Printer is fully configured.
func NewPrinter(conf PrinterConfig) (p *Printer, err error) {
	p = &Printer{}
	// TODO: Probably rewrite two blocks below using unexported funcs.
	p.bold, p.underline, p.reversed, p.blinking = conf.Bold, conf.Underline, conf.Reversed, conf.Blinking
	p.prefix, p.suffix = conf.Prefix, conf.Suffix
	if conf.Color != nil {
		if err = p.SetColor(conf.Color); err != nil {
			return
		}
	}
	if conf.Background != nil {
		if err = p.SetBackground(conf.Background); err != nil {
			return
		}
	}
	err = nil
	return
}

// Printing methods

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

// Color methods

// SetColor sets color of printer. Argument "color" must be either string or int.
// Valid color names are: "black", "blue", "cyan", "green", "magenta",
// "red", "white", "yellow", "brightblack", "brightblue", "brightcyan",
// "brightgreen", "brightmagenta", "brightred", "brightwhite", "brightyellow".
// Valid numbers are from 0 to 255 inclusive.
// If color is not known, is empty, or int is out of range the method
// will return an error and currently set Printer color will not be changed.
func (p *Printer) SetColor(color interface{}) error {
	if code, err := getColorCode(color); err == nil {
		p.color = code
		return nil
	}
	return ErrFailedToSetColor
}

// SetBackground sets background color of printer. Argument color is the same as in SetColor method.
// will return an error and currently set Printer color will not be changed.
func (p *Printer) SetBackground(color interface{}) error {
	if code, err := getBackgroundCode(color); err == nil {
		p.background = code
		return nil
	}
	return ErrFailedToSetBackground
}

// SetPrefixSuffix configures Printer to always preceed output with prefix
// and end output with suffix. Printer color and style settings apply to prefix and suffix.
func (p *Printer) SetPrefixSuffix(prefix, suffix string) {
	p.prefix, p.suffix = prefix, suffix
}

// Modes Methods

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

// ToggleBlinking toggles blinking mode of Printer
func (p *Printer) ToggleBlinking() {
	p.blinking = !p.blinking
}

// Reset resets printer state to initial state (no color, no background, bold, underline and reversed modes turned off).
func (p *Printer) Reset() {
	p.color = ""
	p.background = ""
	p.bold = false
	p.underline = false
	p.reversed = false
	p.blinking = false
	p.prefix = ""
	p.suffix = ""
}

// Methods implementing cursor movement

// PrintAtPosition moves cursor to specified column and row and issues Print
// It does not return to the initial position. It returns the number of bytes written and any write error encountered.
// See also PrintAtPositionAndReturn method.
func (p *Printer) PrintAtPosition(column, row int, a ...interface{}) (n int, err error) {
	out := p.processString(a...)
	moveCursorTo(column, row)
	return fmt.Print(out)
}

// PrintAtPositionAndReturn moves cursor to specified column and row and issues Print
// then moves cursor to initial position when method was called. It returns the number of bytes written and any write error encountered.
func (p *Printer) PrintAtPositionAndReturn(column, row int, a ...interface{}) (n int, err error) {
	out := p.processString(a...)
	saveCursorPosition()
	moveCursorTo(column, row)
	defer restoreCursorPosition()
	return fmt.Print(out)
}

// MoveTo places cursor at the specified column and row.
func (p *Printer) MoveTo(column, row int) {
	moveCursorTo(column, row)
}

// MoveHome places cursor at the top left corner of the screen.
func (p *Printer) MoveHome() {
	moveCursorHome()
}

// MoveUp moves cursor up specified number of rows.
func (p *Printer) MoveUp(rows int) {
	moveCursorUp(rows)
}

// MoveDown moves cursor down specified number of rows.
func (p *Printer) MoveDown(rows int) {
	moveCursorDown(rows)
}

// MoveLeft moves cursor left specified number of columns.
func (p *Printer) MoveLeft(columns int) {
	moveCursorLeft(columns)
}

// MoveRight moves cursor right specified number of columns.
func (p *Printer) MoveRight(columns int) {
	moveCursorRight(columns)
}

// MoveToNextRow moves cursor to the next row..
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
	if p.blinking {
		out += Blinking
	}
	if out == "" {
		return p.prefix + fmt.Sprint(a...) + p.suffix
	}
	out += p.prefix + fmt.Sprint(a...) + p.suffix + Reset
	return out
}
