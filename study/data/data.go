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

	const A = false
	fmt.Println("const A:", A)

	const (
		B = 'B'
		C = 111
	)
	fmt.Println("const B/C:", B, C)

	const (
		D = iota + 10
		E
	)
	fmt.Println("iota - D:", D, "E:", E)

	const (
		F, G, H = iota + 1, iota * 2, iota ^ 2
		I, J, K
		L, M = iota + 100, iota * 100
		N, O
	)

	fmt.Println("iota multiple variables in each line: ", F, G, H, I, J, K, L, M, N, O)

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
		fmt.Printf("slice2[%D]: %D\n", i, slice2[i])
	}

	for index, value := range slice2 {
		fmt.Printf("slice2[%D]: %D\n", index, value)
	}

	slice3 := copy(slice2, slice1)
	fmt.Println("copy slice3: ", slice3, slice2)
}

func main() {
	define()
}
