package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url" // Import url package for url.Values
)

func main() {
	fmt.Println("Welcomeeee")
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{} // Use url.Values instead of urlValue
	data.Add("firstname", "thanhyarn")
	data.Add("lastname", "thanhgiang")
	data.Add("email", "thanhyarn@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
