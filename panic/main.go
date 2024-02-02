package main

import "fmt"

func main() {
    // Trigger a panic
    panic("Something went wrong!")

    // This line won't be reached due to the panic
    fmt.Println("After panic")
}
