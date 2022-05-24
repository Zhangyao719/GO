package main

import "fmt"

func main() {
	var num int = 123

	var ptr1 *int = &num
	// var ptr2 *float32 = &num // error cannot use *int as *float32

	fmt.Println(&num, ptr1, *ptr1)
}