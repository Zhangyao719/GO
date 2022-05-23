package main

import (
	"fmt"
	"strconv"
)

func main() {
	n1 := 999
	n2 := 23.456
	b := true
	var chart byte = 'h'
	var str string

	// * 方式一：fmt.Sprintf
	// * 参数一： 格式，参数二：数值
	str = fmt.Sprintf("%d", n1) // %d 十进制
	fmt.Printf("str的类型:%T, str的值:%q\n", str, str) // %q 单引号围绕的字符字面值

	str = fmt.Sprintf("%f", n2) // %d 浮点
	fmt.Printf("str的类型:%T, str的值:%q\n", str, str)

	str = fmt.Sprintf("%t", b) // %t 布尔值
	fmt.Printf("str的类型:%T, str的值:%q\n", str, str)

	str = fmt.Sprintf("%c", chart) // %c 字符
	fmt.Printf("str的类型:%T, str的值:%q\n", str, str)

	// * 方式二： strconv.Format
	str = strconv.FormatInt(int64(n1), 10)
	fmt.Printf("str的类型:%T, str的值:%q\n", str, str)

	str = strconv.FormatFloat(n2, 'f', 10, 64)
	fmt.Printf("str的类型:%T, str的值:%q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str的类型:%T, str的值:%q\n", str, str)

	str = strconv.Itoa(n1)
	fmt.Printf("str的类型:%T, str的值:%q\n", str, str)

}