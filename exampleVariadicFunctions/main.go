package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    // Call the sum function with different numbers of arguments
    result1 := sum(1, 2, 3, 4, 5)
    result2 := sum(10, 20, 30)
    result3 := sum(5, 10, 15, 20, 25, 30)

    // Print the results
    fmt.Println("Result 1:", result1)
    fmt.Println("Result 2:", result2)
    fmt.Println("Result 3:", result3)
}
