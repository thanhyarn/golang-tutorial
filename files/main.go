package main

import (
	"fmt"
	"io"
	"io/ioutil" // Import the ioutil package
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "Content ABC DEF"

	file, err := os.Create("./test.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./test.txt") // Correct the filename
}

func readFile(fileName string) { // Fix the variable name
	datatype, err := ioutil.ReadFile(fileName) // Fix the function call
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(datatype))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
