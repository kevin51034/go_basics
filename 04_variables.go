package main


import "fmt"


var i int
var python, java bool
var golang,c = true, "c!"

func main() {
	//golang = true
	// Outside a function, every statement begins with a keyword (var, func, and so on) 
	// and so the := construct is not available, cann't use ouside the function
	k := 3
	fmt.Println(i,k,golang,c,python,java)
}