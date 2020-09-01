package main

import (
	"fmt"

	"github.com/dmfed/termtools"
)

func main() {
	a := "hello, world"
	fmt.Println(termtools.Paint("red", a)
	fmt.Printf("%vHello, world!%v\n", termtools.Blue, termtools.ColorReset)
	fmt.Printf("%vWe are gophers.%v\n", termtools.ByID(196), termtools.Reset())
	fmt.Println(termtools.Paint("red", "We like Go"))
}
