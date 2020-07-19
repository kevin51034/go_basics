package main

import "fmt"

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v:=range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}