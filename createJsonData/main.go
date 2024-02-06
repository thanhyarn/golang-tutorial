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
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{ // Corrected the type to 'course'
		{"ReactJS Bootcamp", 299, "thanhyarn.dev", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "thanhyarn.dev", "gc123", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 399, "thanhyarn.dev", "ackac123", nil},
	}

	// package this data as JSON data

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
