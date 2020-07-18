package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// slice is a dynamically-sized, flexible view into the elements of an array.
	var s []int = primes[1:4] // from index 1 to 4, 4 is not included
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// ** Slices are like references to arrays
	// ** Changing the elements of a slice modifies the corresponding elements of its underlying array.
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// slice literal is like an array literal without the length.
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	fmt.Println(len(q))

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	k := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(k)


	h := []int{2, 3, 5, 7, 11, 13}

	h = h[1:4]
	fmt.Println(h)

	h = h[:2]
	fmt.Println(h)

	h = h[1:]
	fmt.Println(h)

}
