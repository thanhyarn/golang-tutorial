package main

import "fmt"

func main(){
	// Khai bái và khởi tạo một map
	person := map[string]int{
		"Nam": 25,
		"Giang": 23,
		"Thanh": 35,
	}

	// In gúa trị của map
	fmt.Println("Map person: ", person)

	// Truy cập giá trị thông qua key
	ageGiang := person["Giang"]
	fmt.Println("Tuổi của Giang là :" , ageGiang)

	// Thêm một cặp key-value vào map 
	person["Nhi"] = 28

	// In giá trị sau khi thêm 
	fmt.Println("Map person sau khi thêm: " , person)
	_, isExits := person["Nam"]
	fmt.Println("Nam tồn tại trong map: ", isExits)

	//Xóa một keu khỏi map
	delete(person, "Thanh")

	//In giá trị sau khi xóa 
	fmt.Println("Map person sau khi xóa: ", person)
	
}