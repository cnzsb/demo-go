package main

import "fmt"

func f1() string {
	fmt.Println("f1() is called")

	return "f1"
}

func f2() (string, string) {
	return "r1", "r2"
}

func f3() (r1 string, r2 string) {
	r1 = "r1"
	r2 = "r2"

	return r1, r2
}

func f4() (r1, r2 string) {
	r1 = "r1"
	r2 = "r2"

	return r1, r2
}

func deferF1() string {
	defer fmt.Println("defer is called")

	return f1()
}

func main() {
	fmt.Println(f1())

	f2r1, f2r2 := f2()
	fmt.Println(f2r1, f2r2)

	f3r1, f3r2 := f3()
	fmt.Println(f3r1, f3r2)

	fmt.Println(f4())

	fmt.Println("defer func: ", deferF1())
}
