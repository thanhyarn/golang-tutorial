package main

import "fmt"

func factorial(n int) int {
    // Base case: factorial of 0 is 1
    if n == 0 {
        return 1
    }
    
    // Recursive case: n! = n * (n-1)!
    return n * factorial(n-1)
}

func main() {
    // Calculate the factorial of 5
    result := factorial(5)

    // Print the result
    fmt.Println("Factorial of 5:", result)
}
