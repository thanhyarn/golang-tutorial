package main

import "fmt"

func main() {
    // Recover from a panic
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()

    // Trigger a panic
    panic("Something went wrong!")

    // This line won't be reached due to the panic
    fmt.Println("After panic")
}
