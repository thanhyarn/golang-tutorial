package main

import "fmt"

func main() {
    // Declare a variable
    var num int = 42

    // Declare a pointer variable
    var ptr *int

    // Assign the address of the variable to the pointer
    ptr = &num

    // Print the value and memory address of the variable
    fmt.Printf("Value of num: %d\n", num)
    fmt.Printf("Memory address of num: %p\n", &num)

    // Print the value and memory address stored in the pointer
    fmt.Printf("Value pointed to by ptr: %d\n", *ptr)
    fmt.Printf("Memory address stored in ptr: %p\n", ptr)

    // Modify the value indirectly through the pointer
    *ptr = 99

    // Print the modified value of the variable
    fmt.Printf("Modified value of num: %d\n", num)
}