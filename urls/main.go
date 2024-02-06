package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// Parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for key, val := range qparams {
		fmt.Printf("Param is: %s - %s\n", key, val[0])
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawPath:  "user=hitech",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
