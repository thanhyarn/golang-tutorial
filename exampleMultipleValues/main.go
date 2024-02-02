package main

import "fmt"

func findMinMax(numbers []int) (int, int) {
    // Check if the slice is empty
    if len(numbers) == 0 {
        fmt.Println("The slice is empty.")
        return 0, 0
    }

    // Initialize min and max with the first element
    min, max := numbers[0], numbers[0]

    // Iterate through the slice to find min and max
    for _, num := range numbers {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }

    return min, max
}

func main() {
    // Declare and initialize a slice
    numbers := []int{48, 96, 86, 68, 57, 82, 63, 70, 19, 97, 9, 17}

    // Call the function to find min and max
    min, max := findMinMax(numbers)

    // Print the results
    fmt.Println("Minimum:", min)
    fmt.Println("Maximum:", max)
}
