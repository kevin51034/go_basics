package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// byte is alias for uint8

// rune is alias for int32
// represents a Unicode code point

func main() {
	// %T a Go-syntax representation of the type of the value
	// %v the value in a default format
	// when printing structs, the plus flag (%+v) adds field names
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	// Constants cannot be declared using the := syntax.
	x, y := 3, 4
	// must use type float64 in argument to math.Sqrt
	s := math.Sqrt(float64(x*x + y*y))
	c := int(s)
	fmt.Println(x, y, s, c)

}