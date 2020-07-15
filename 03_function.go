package main

import "fmt"

func main() {
	fmt.Println(adder(16,43))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(10))
}

func adder(x int,y int) int{
	return x+y
}

// multiple return value

func swap(x string, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 10
	y = sum / 2
	return
	// A return statement without arguments returns the named return values.
	// This is known as a "naked" return.
}