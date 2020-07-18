package main

import "fmt"

type Person struct {
	name string
	age int
}


func main() {
	fmt.Println(Person{"kevin", 22})
	person2 := Person{"john", 30}
	fmt.Println(person2.name, person2.age)

	// pointer to struct
	p := &person2
	p.age = p.age + 1
	fmt.Println(person2.name, person2.age)


	// struct literals
	var (
		p1 = Person{"p1", 20}  // has type Person
		p2 = Person{name:"p2"}  // age is implicit
		p3 = Person{}      // name:"" and age:0
		g  = &Person{"p1", 20} // has type *Person
	)
	
	fmt.Println(p1, g, p2, p3)
	
}