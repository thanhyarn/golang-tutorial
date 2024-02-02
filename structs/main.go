package main

import "fmt"

// Define a struct named "Person"
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    // Create an instance of the Person struct
    person1 := Person{
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
    }

    // Access and print the struct fields
    fmt.Println("First Name:", person1.FirstName)
    fmt.Println("Last Name:", person1.LastName)
    fmt.Println("Age:", person1.Age)

    // Create another instance of the Person struct
    person2 := Person{
        FirstName: "Jane",
        LastName:  "Doe",
        Age:       25,
    }

    // Access and print the struct fields of the second person
    fmt.Println("\nSecond Person:")
    fmt.Println("First Name:", person2.FirstName)
    fmt.Println("Last Name:", person2.LastName)
    fmt.Println("Age:", person2.Age)

    // Modify a field of the first person
    person1.Age = 31

    // Print the updated information
    fmt.Println("\nUpdated First Person:")
    fmt.Println("First Name:", person1.FirstName)
    fmt.Println("Last Name:", person1.LastName)
    fmt.Println("Age:", person1.Age)
}
