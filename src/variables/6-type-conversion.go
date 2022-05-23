package main

import "fmt"

func main() {
	var i int32 = 100
	// * 类型转换
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 uint64 = uint64(i)

	fmt.Printf("n1: %T %v \n n2: %T %v \n n3: %T %v \n", n1, n1, n2, n2, n3, n3)

	// * 在转换中，如果 高精度 超出 低精度 范围，编译并不会报错，而会做溢出处理。
	// * 比如 uint64 999 转成 uint8 (999 超出了 0-255 范围)
	var n4 uint64 = 999
	var n5 uint8 = uint8(n4)
	fmt.Println("n5", n5)
	
	var n6 int32 = 300
	var n7 int8 = int8(n6) + 127 // 编译通过 溢出处理
	// var n8 int8 = int8(n6) + 128 // error 128本身已超出 int8 范围
	fmt.Println("n7", n7)

}