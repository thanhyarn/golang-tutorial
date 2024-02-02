package main

import "fmt"

func main(){
	// The function foo will me deferred until the end of the main function
	defer foo()

	fmt.Println("Hello")
}

func foo(){
	fmt.Println("Deferred function: foo")
}