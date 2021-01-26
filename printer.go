package termtools

import (
	"fmt"
)

// Printer implements most methods from fmt like Print, Println, Sprint
// adding colors and styles to input values
type Printer struct {
	color      string
	background string
	bold       bool
	underline  bool
	reversed   bool
}

// NewColorPrinter takes color name (string) and returns Printer
// with font color set to the requirement. If supplied color name
// is invalid, the function will return an error.
func NewColorPrinter(color string) (*Printer, error) {
	var p Printer
	code, err := GetColorByName(color)
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

func (p *Printer) Println(a ...interface{}) (n int, err error) {
	out := p.processString(a...)
	return fmt.Println(out)
}

func (p *Printer) Print(a ...interface{}) (n int, err error) {
	out := p.processString(a...)
	return fmt.Print(out)
}

func (p *Printer) SetColor(colorname string) error {
	code, err := getColorByName(colorname)
	if err == nil {
		p.color = code
	}
	return err
}

func (p *Printer) SetColorID(id int) error {
	code, err := getColorByID(id)
	if err == nil {
		p.color = code
	}
	return err
}

func (p *Printer) SetBackground(colorname string) error {
	code, err := getBackgroundByName(colorname)
	if err == nil {
		p.background = code
	}
	return err
}

func (p *Printer) SetBackgroundID(id int) error {
	code, err := getBackgroundByID(id)
	if err == nil {
		p.background = code
	}
	return err
}

func (p *Printer) ToggleBold() {
	p.bold = !p.bold
}

func (p *Printer) ToggleUnderline() {
	p.underline = !p.underline
}

func (p *Printer) ToggleReversed() {
	p.reversed = !p.reversed
}

func (p *Printer) Reset() {
	p.color = ""
	p.background = ""
	p.bold = false
	p.underline = false
	p.reversed = false
}

/*

func Printf(format string, a ...interface{}) (n int, err error) {

}

func Sprint(a ...interface{}) string {

}

func Sprintf(format string, a ...interface{}) string {

}

func Sprintln(a ...interface{}) string {

}
*/

/*
func Errorf(format string, a ...interface{}) error                            {}
func Fprint(w io.Writer, a ...interface{}) (n int, err error)                 {}
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {}
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)               {}
func Fscan(r io.Reader, a ...interface{}) (n int, err error)                  {}
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)  {}
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)                {}
                               {}

func Scan(a ...interface{}) (n int, err error)                                {}
func Scanf(format string, a ...interface{}) (n int, err error)                {}
func Scanln(a ...interface{}) (n int, err error)                              {}
                             {}
func Sscan(str string, a ...interface{}) (n int, err error)                   {}
func Sscanf(str string, format string, a ...interface{}) (n int, err error)   {}
func Sscanln(str string, a ...interface{}) (n int, err error)                 {}
*/

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
