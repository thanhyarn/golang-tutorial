package main

import (
	"encoding/json" // Add the json package
	"fmt"
)

// type course struct {
// 	Name     string`
// 	Print    int
// 	Platform string`
// 	Password string`
// 	Tags     []string`		
// }

type course struct {
	Name     string `json:"coursename"`
	Print    int
	Platform string `json:"website"`
	Password string `json: "-"`
	Tags     []string `json:"tags, omitempty"`
}

func main() {
	fmt.Println("Welcome JSON")
	//EncodeJson()
	DecodeJson()
}

func DecodeJson(){
	jsonDataFormWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "yarndev.wordpress.com",
			"tags": ["web-dev", "js"]
		}	
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFormWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFormWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// var cases where you just want to add data to key value
	
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFormWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData{
		fmt. Printf("Key is %v and value is %v and Type is: %T\n", k,v,v)
	}
}