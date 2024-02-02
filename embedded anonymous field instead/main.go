package main

import "fmt"

// Animal type with methods
type Animal struct {
    Name string
}

func (a Animal) MakeSound() {
    fmt.Println("Generic animal sound")
}

// Dog type embedding Animal
type Dog struct {
    Animal
    Breed string
}

func (d Dog) MakeSound() {
    fmt.Println("Bark! Bark!")
}

func main() {
    // Using embedded anonymous field
    dog := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden Retriever"}
    fmt.Println("Dog's name:", dog.Name)
    fmt.Println("Dog's breed:", dog.Breed)
    dog.MakeSound() // Calls the overridden method in Dog

    // Using normal named field
    type Cat struct {
        Animal   Animal
        Breed    string
        CollarID string
    }
    cat := Cat{Animal: Animal{Name: "Whiskers"}, Breed: "Siamese", CollarID: "12345"}
    fmt.Println("Cat's name:", cat.Animal.Name)
    fmt.Println("Cat's breed:", cat.Breed)
    cat.Animal.MakeSound() // Calls the method in Animal
}
