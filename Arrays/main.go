package main

import "fmt"

func main()  {
	//Khai báo một mảng có 5 phần tử kiểu int
	var Arr[5]int

	// Gán giá trị cho các phần tử trong mảng
	Arr[0] = 10;
	Arr[1] = 20;
	Arr[2] = 10;
	Arr[3] = 20;
	Arr[4] = 10;

	// Hoặc có thể khỏi tạo mảng ngay từ đầu
	// Arr :=[5]int{10, 20, 30, 40, 60}

	// In các phần tử của mảng
    fmt.Println(Arr)
	
}