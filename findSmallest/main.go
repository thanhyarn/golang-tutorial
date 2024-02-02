package main

import "fmt"

func main() {
    // Declare and initialize an array
    x := []int{
        48, 96, 86, 68,
        57, 82, 63, 70,
        19, 97, 9, 17,
    }

    // Find the smallest number
    smallest := findSmallest(x)

    // Print the result
    fmt.Println("The smallest number is:", smallest)
}

// Function to find the smallest number in a slice
func findSmallest(numbers []int) int {
    // Check if the slice is empty
    if len(numbers) == 0 {
        fmt.Println("The slice is empty.")
        return 0
    }

    // Initialize the smallest with the first element
    smallest := numbers[0]

    // Iterate through the slice to find the smallest number
    for _, num := range numbers {
        if num < smallest {
            smallest = num
        }
    }

    return smallest
}
