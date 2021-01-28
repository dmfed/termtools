package termtools

import (
	"fmt"
	"io"
)

// Printer holds color and style settings and implements most methods as in fmt modure like Print, Println, Sprint etc.
// adding color and styles to the input values.
type Printer struct {
	// contains unexported fields
	color      string
	background string
	bold       bool
	underline  bool
	reversed   bool
}

// NewColorPrinter takes color name (string) and returns Printer with font color set to the requirement.
// If supplied color name is invalid, the function will return an error.
func NewColorPrinter(color string) (*Printer, error) {
	var p Printer
	code, err := getColorByName(color)
	p.color = code
	return &p, err
}

// NewPrinter returns instance of Printer with parameters set as requested
func NewPrinter(color, background string, bold, underline, reversed bool) (*Printer, error) {
	p, err := NewColorPrinter(color)
	p.bold = bold
	p.underline = underline
	p.reversed = reversed
	return p, err
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
func (p *Printer) SetColorID(id int) error {
	code, err := getColorByID(id)
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
func (p *Printer) SetBackgroundID(id int) error {
	code, err := getBackgroundByID(id)
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
	out += fmt.Sprint(a...)
	out += Reset
	return out
}
