package main

import (
	tt "github.com/dmfed/termtools"
)

var prnt = tt.PrintSuite{}

func main() {
	prnt.Println("Default mode. No settings.")
	p1, p2 := tt.Printer{}, tt.Printer{}
	p1.SetColor("green")
	p2.SetColor("blue")
	prnt.AddPrinter("green", &p1)
	prnt.AddPrinter("blueblink", &p2)
	prnt.Use("green").Println("This is green")
	prnt.SwitchTo("blueblink")
	prnt.Println("I like blue.")
	prnt.Reset()
	prnt.Println("Reset blinking mode with Reset()")
	prnt.SwitchTo("blueblink")
	prnt.Println("Restored blue blinking configuration with SwitchTo")
	prnt.AddPrinter("nil", nil)
	prnt.Use("nil").Println("This is a call to Use(\"nil\").Println")
}
