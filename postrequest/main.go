package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("Welcomeeee")
	PerformPostRequest()
}

func PerformPostRequest(){
	const myurl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": "20",
			"platform": "yarndev.wordpress.com"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil{
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

