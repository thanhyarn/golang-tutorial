package main

import "fmt"

func main() {
    // Outer function that returns a closure
    outer := outerFunction(10)

    // Call the closure
    result := outer(5)

    // Print the result
    fmt.Println("Result:", result)
}

// Outer function that returns a closure
func outerFunction(x int) func(y int) int {
    // The inner function is a closure that references the variable 'x'
    inner := func(y int) int {
        return x + y
    }

    // Return the closure
    return inner
}
