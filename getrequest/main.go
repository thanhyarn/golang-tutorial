package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings" // Add import for the strings package
)

func main() {
	fmt.Println("Welcome to Get Request")
	PerformGetRequest() // Call the PerformGetRequest function
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())
	fmt.Println(content)
	fmt.Println(string(content))
}
