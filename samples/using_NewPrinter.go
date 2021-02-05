package main

import "github.com/dmfed/termtools"

func main() {
	printer, _ := termtools.NewPrinter(termtools.PrinterConfig{Color: "blue", Blinking: true})
	printer.Println("Printing in blue blinking here.")
}
