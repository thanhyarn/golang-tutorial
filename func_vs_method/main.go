package main

import "fmt"

// Define a type named 'Person'
type Person struct {
    FirstName string
    LastName  string
}

// Function to get the full name of a person
func getFullName(person Person) string {
    return person.FirstName + " " + person.LastName
}

// Method to get the full name of a person (associated with the 'Person' type)
func (p Person) getFullNameMethod() string {
    return p.FirstName + " " + p.LastName
}

func main() {
    // Create an instance of the 'Person' type
    johnDoe := Person{"John", "Doe"}

    // Using the function
    fullNameFunction := getFullName(johnDoe)
    fmt.Println("Using function:", fullNameFunction)

    // Using the method
    fullNameMethod := johnDoe.getFullNameMethod()
    fmt.Println("Using method:", fullNameMethod)
}
