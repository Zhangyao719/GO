package main

import "fmt"

func main() {
	// * 十进制：
	num := 3.14

	// * 科学计数法
	num1 := 3.14e2  // 314
	num2 := 3.14E2  // 314
	num3 := 3.14E-2 // 0.0314

	fmt.Println(num, num1, num2, num3)
}
