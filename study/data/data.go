package main

import "fmt"

func define() {
	var str1 string
	str1 = "hello"
	str2 := "go"
	fmt.Println("string:", str1, str2)

	// default value is 0
	var num1 int
	num2 := 3
	fmt.Println("num:", num1, num2)

	//const (
	//	IOTA1 = iota + 100
	//	IOTA2,
	//	IOTA3, IOTA4 = iota * 2, iota * 3
	//	IOTA5, IOTA6,
	//	IOTA7, IOTA8, IOTA9 = iota + 10, iota + 20, iota + 30
	//	IOTA10, IOTA11, IOTA12
	//)
	//fmt.Println("IOTA: ", IOTA1, IOTA2, IOTA3, IOTA4, IOTA5, IOTA6, IOTA7, IOTA8, IOTA9, IOTA10, IOTA11, IOTA12)

	var arr1 [3]int
	arr1[1] = 1

	arr2 := [5]int{1}

	fmt.Println("arr values: ", arr1, arr2)
	fmt.Println("arr len/cap: ", len(arr1), cap(arr1), len(arr2), cap(arr2))

	// dynamic array is slice
	var slice1 []int
	slice1 = append(slice1, 1)
	fmt.Println("slice1: ", slice1, len(slice1), cap(slice1))

	slice2 := []int{1, 2, 3}
	fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 4)
	fmt.Println("slice2 after appended a new element: ", slice2, len(slice2), cap(slice2))

	for i := 0; i < len(slice2); i++ {
		fmt.Printf("slice2[%d]: %d\n", i, slice2[i])
	}

	for index, value := range slice2 {
		fmt.Printf("slice2[%d]: %d\n", index, value)
	}

	slice3 := copy(slice2, slice1)
	fmt.Println("copy slice3: ", slice3, slice2)
}

func main() {
	define()
}
