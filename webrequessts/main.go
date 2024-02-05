package main

import (
	"fmt"
	"io/ioutil"
	"net/http" // Import the net/http package
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	datatypes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(datatypes)
	fmt.Println(content)
}
