package main

import "github.com/dmfed/termtools"

var prnt = termtools.PrintSuite{} // we define prnt globally to reuse it from any part of program

func setupPrinters() {
	configs := []termtools.PrinterConfig{
		{Name: "error", Color: "red", Prefix: ">>> "},
		{Name: "notify", Color: "green", Underline: true}}
	if err := prnt.Configure(configs...); err != nil {
		prnt.Println(err)
	}
}

func main() {
	setupPrinters()
	// We can use different printers for different kinds of messages anywhere in out program now.
	// For example print errors in red
	prnt.Use("error").Print("We call prnt.Use(\"error\").Print() to output using printer named \"error\"\n")
	// or print notifications in green
	prnt.Use("notify").Print("and prnt.Use(\"notify\").Print() to use printer named \"notify\"\n")
	// or use unstyled output for everything else like this:
	prnt.Println("This will print without any styling.")
	// or like this:
	prnt.UseDefault().Println("This will print without any styling.")
}
