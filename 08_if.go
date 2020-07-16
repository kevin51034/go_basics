package main

import(
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	// Notice the order would be pow1(return) -> pow2(return) -> fmt.Println
	// in this case output is 27>20\n 9 20
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func sqrt(x float64) string {
	if x<0 {
		return sqrt(-x) + "i"
	}
	// Sprint formats using the default formats for its operands and returns the resulting string.
	// Spaces are added between operands when neither is a string.
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// Like for, the if statement can start with a short statement to execute before the condition.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}