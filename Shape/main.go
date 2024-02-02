package main

import "fmt"

// Shape interface with methods for area and perimeter
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Circle type implementing the Shape interface
type Circle struct {
    Radius float64
}

// Rectangle type implementing the Shape interface
type Rectangle struct {
    Width  float64
    Height float64
}


// Implementing the Area method for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Implementing the Perimeter method for Circle
func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

// Implementing the Area method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Implementing the Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
    return 2*r.Width + 2*r.Height
}

func main() {
    // Creating instances of Circle and Rectangle
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 4, Height: 6}

    // Using the Area and Perimeter methods through the Shape interface
    printShapeInfo(circle)
    printShapeInfo(rectangle)
}

// Function to print area and perimeter of a shape
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}
