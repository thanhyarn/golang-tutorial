package main 

import "fmt"

func main(){
	//Khai báo slice 
	var sliceArr []int

	// Hoặc sử dụng make để tạo ra slice với kích thước cho trước
	// sliceArr := make([]int, 5, 10)

	// Thêm phần tử vào slice 
	sliceArr = append(sliceArr, 10 , 20 , 35, 60, 75, 80)

	// In các phần tử của slice
	fmt.Println(sliceArr)


	// Thay đổi giá trị của một phần tử trong slice
	sliceArr[2] = 123;
	fmt.Println("Slice sau khi thay đổi giá trị", sliceArr)

	// Tạo một slice con từ một slice có sẵn
	currentSlice := sliceArr[1:4]
	fmt.Println("Current Slice: " , currentSlice)

	// Lưu ý : Slices chia sẻ cùng một mảng dữ liệu
	slice1 := sliceArr[1:3]
	slice2 := sliceArr[2:4]

	// Thay đổi một phần tử trong slice1
	slice1[0] = 100

	// In giá trị của cả hai mảng slice
	fmt.Println("Slice 1: ", slice1)
	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Slice gốc sau khi thay đổi ", sliceArr )
}