package main

import "github.com/dmfed/termtools"

var prnt = termtools.PrintSuite{}

var configAddCode string = `configs := []termtools.PrinterConfig{
	{Name: "error", Color: "red"},
	{Name: "notify", Color: "green", Underline: true}}
if err := prnt.Configure(configs...); err != nil {
	prnt.Println(err)
}`

func main() {
	termtools.ClearScreen()
	termtools.MoveCursorHome()
	prnt.Println("This is an example of using PrintSuite.\n\nWe just declared var prnt = termtools.PrintSuite{}. This is the default embedded printer printing\nusing prnt.Println(). It is safe to declare PrintSuite and use it straight away.\n")
	configs := []termtools.PrinterConfig{
		{Name: "error", Color: "red"},
		{Name: "notify", Color: "green", Underline: true}}
	if err := prnt.Configure(configs...); err != nil {
		prnt.Println(err)
	}
	prnt.Println("Added two configurations named \"error\" and \"notify\" using this piece of code:\n")
	prnt.Println(configAddCode, "\n")
	prnt.SwitchTo("notify")
	prnt.Println("Now switched default embedded printer configuration using SwitchTo(\"notify\"). If we just call\nprnt.Println() we will continue printing in green.\n")
	prnt.Use("error").Println("This line is printed using prnt.Use(\"error\").Println()\n")
	prnt.UseDefault().Println("Printing with prnt.UseDefault().Println() here so will print without any styles set.\n")
	prnt.Println("Calling prnt.Println() Note that the embedded printer has not changed after call to\nprnt.SwitchTo(\"notify\") above.\n")
	prnt.SwitchToDefault()
	prnt.Println("Now we called prnt.SwitchToDefault() and printing this with prnt.Println()\nNote that the embedded printer now has no style settings.\n")
	prnt.Use("error").Print("We can still call prnt.Use(\"error\").Print() ")
	prnt.Use("notify").Print("and prnt.Use(\"notify\").Print()\n")
}
