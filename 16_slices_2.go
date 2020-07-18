package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	var k []int
	fmt.Println(k, len(k), cap(k))
	if k == nil {
		fmt.Println("nil!")
	}


	// creating a slice with make
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)


	// slices of slices
	// slices can contain any type, including other slices.

		// Create a tic-tac-toe board.
		board := [][]string{
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
		}
	
		// The players take turns.
		board[0][0] = "X"
		board[2][2] = "O"
		board[1][2] = "X"
		board[1][0] = "O"
		board[0][2] = "X"
	
		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
		}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}