package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct{
	CourseId 	string	`json: "courseid"`
	CourseName	string 	`json: "coursename"`
	CoursePrice int		`json: "price"`
	Author 		*Author	`json: "author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website	 string `json:"website"`
}

var courses []Course

//middleware , helper - file
func (c *Course) IsEmpty() bool{
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main(){

}

//controller - file 

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get on course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars()

	//loop through courses, find matching id and return the response
	for _, course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return 
}

func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create on course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id , string 
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	return 
}

