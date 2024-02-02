package main 

import "fmt"

func main(){

	// HOW TO YOU ACCESS THE 4TH ELEMENT OF AN ARRAY OR SLICE ?

	// Declare and initialize a slice 
	number := []int{10, 20 , 30 , 40, 50}

	//Access the fourth element (index 3)
	fourthElement := numbers[3];

	//Print the value of the fourth element
	fmt.Println("Fourth element : ", fourthElement)	

	//WHAT IS THE LENGTH OF A SLICE CREATED USING MAKE([]INT , 3 , 9)

	mySlice := make([]int, 3, 9);
	
	//Assign values to the element of the slice
	mySlice[0] = 10
	mySlice[1] = 20
	mySlice[2] = 30

	// Print the elements , length, and capacity fo the slice
	fmt.Println("Slice elements:", mySlice)
    fmt.Println("Slice length:", len(mySlice))
    fmt.Println("Slice capacity:", cap(mySlice))

	//Give the array 
	//x := [6]string{"a","b","c","d","e","f"}
	//what would x[2:5] give you?

	// Declare and initialize an array
    x := [6]string{"a", "b", "c", "d", "e", "f"}

    // Create a slice using the array
    y := x[2:5]

    // Print the slice
    fmt.Println("Slice y:", y)


}