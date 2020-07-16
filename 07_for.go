package main

import "fmt"

func main() {
	sum:=0
	for i:=0 ; i<10; i++ {
		sum+=i
	}
	fmt.Println(sum)
	
	// The init and post statements are optional.
	sum=1
	for ; sum < 1000 ; {
		sum+=sum
	}
	fmt.Println(sum)

	// For is Go's "while"
	sum=1
	for sum < 1000 {
		sum+=sum
	}
	fmt.Println(sum)
}