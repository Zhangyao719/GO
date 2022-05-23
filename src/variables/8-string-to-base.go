package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, n1, n2, n3 := "true", "-256", "256", "3.1415"
	params1, _ := strconv.ParseBool(b)
	params2, _ := strconv.ParseInt(n1, 10, 32)
	params3, _ := strconv.ParseUint(n2, 10, 32)
	params4, _ := strconv.ParseFloat(n3, 64)

	fmt.Printf("params1 类型：%T 值：%v \n", params1, params1)
	fmt.Printf("params2 类型：%T 值：%v \n", params2, params2)
	fmt.Printf("params3 类型：%T 值：%v \n", params3, params3)
	fmt.Printf("params4 类型：%T 值：%v \n", params4, params4)

	// * 转换不成功 将变成默认值
	params5, _ := strconv.ParseBool("hello")
	params6, _ := strconv.ParseInt("hello", 10, 32)
	fmt.Printf("params5 类型：%T 值：%v \n", params5, params5) // bool: false
	fmt.Printf("params6 类型：%T 值：%v \n", params6, params6) // int64: 0
}