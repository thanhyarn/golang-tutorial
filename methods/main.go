package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	// Create a User instance with all fields initialized
	thanhgiang := User{
		Name:   "ThanhYarn",
		Email:  "thanhgiang.user@gmail.com",
		Status: true,
		Age:    25,
	}

	fmt.Printf("Thanh Giang details are: %+v\n", thanhgiang)
	fmt.Printf("Name is %v, email is %v, and age is %v.\n", thanhgiang.Name, thanhgiang.Email, thanhgiang.Age)

	// Call the GetStatus method on the User instance
	thanhgiang.GetStatus()
}

type User struct {
	Name  string
	Email string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}
