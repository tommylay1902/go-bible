package main

import (
	"fmt"

	"github.com/tommylay1902/go-bible/flag"
	newkeyword "github.com/tommylay1902/go-bible/newKeyword"
	"github.com/tommylay1902/go-bible/tempconv"
)

func main() {
	flag.FlagFunction()
	newkeyword.NewKeyword()

	var c tempconv.Celsius = 36.77778
	var f tempconv.Fahrenheit = 98.2

	convertedC := tempconv.CToF(c)
	convertedF := tempconv.FToC(f)

	fmt.Printf("%vF converted to c: %v\n", f, convertedF)
	fmt.Printf("%vC converted to f: %v\n", c, convertedC)

	// this will throw a compile error since f is of type tempconv.Fahrenheit
	// and not tempconv.Celsius
	// convertedCFail := tempconv.CToF(f)

	// same error but vice versa
	// convertedFFail := tempconv.FToC(c)

}
