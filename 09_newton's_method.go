package main

import(
	"fmt"
	"math"
)

func main() {
    guess := sqrt(2)
    expected := math.Sqrt(2)
	fmt.Printf("Guess: %v, Expected: %v, Error: %v\n", guess, expected, math.Abs(guess - expected))
}

func sqrt(x float64) float64 {
	z:=float64(1)
	var t float64
	for i:=0;i<10;i++ {
		z, t = z - (z*z - x) / (2*z), z
		fmt.Println(z)
		if math.Abs(t-z) < 1e-6 {
            break
        }
	}
	return z
}

