package main

import "fmt"

// 批量声明方式一：
var global1, global2, global3 = 1, 2, 3

// 批量声明方式二：
var (
	global4 = 4
	global5 = 5
	global6 = 6
)

func main() {
	// * 第一种
	var num0 int = 10

	// * 第二种 声明类型 不赋值（使用默认值）
	var num1 int
	
	// * 第三种 直接赋值（类型推导）
	var num2 = 10.22

	// * 第四种 :=
	name := "Tom"

	// * 批量声明
	var num3, num4, num5 int
	var num6, num7, num8 = 600, "700", 800 // num6:600 num7:"700" num8:800
	// num6, num7, num8 := 600, "700", 800

	fmt.Println(num0, num1, num2, name, num3, num4, num5, num6, num7, num8)

}
