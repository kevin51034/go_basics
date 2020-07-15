package main

import "fmt"

func main() {
	fmt.Println(adder(16,43))
	fmt.Println(swap("hello", "world"))
}

func adder(x int,y int) int{
	return x+y
}

// multiple return value

func swap(x string, y string) (string, string) {
	return y, x
}