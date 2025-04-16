package main

import (
	"fmt"
)

func NotMain() {
	fmt.Print("TypedAndUntypedConsts")
	const x = 10      // untyped constant
	const xx int = 10 // typed constant, можно присвоить только переменной типа int

	// xx = 20 // error: cannot assign to xx (xx is a constant)

	var y int = x     // x is a constant of type int
	var z float64 = x // x is a constant of type float64
	var d byte = x    // x is a constant of type byte

	fmt.Print(x, xx, y, z, d)
}
