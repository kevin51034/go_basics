package main

import "fmt"


func main() {
	i, j := 1, 0

	p:= &i
	fmt.Println(*p) //1
	*p=21
	fmt.Println(i) //21
	p=&j
	fmt.Println(*p) //0
	*p = *p + 100
	fmt.Println(j) //100

	// Unlike C, Go has no pointer arithmetic.


}