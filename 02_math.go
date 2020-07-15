package main

import(
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n",
				math.Nextafter(2,3))
	// returns the next representable value after x towards y.
	// which is Float64frombits(Float64bits(x) + 1)

	fmt.Println(math.Pi)
	// math.pi would be undefined
	// In Go, a name is exported if it begins with a capital letter.
}