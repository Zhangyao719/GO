package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int64 = 12

	// num → %T，unsafe.Sizeof(num) → %d
	fmt.Printf("num 类型 %T 字节数 %d", num, unsafe.Sizeof(num))
}

// * 使用整数型变量时，遵守保大不保小的原则，尽量使用占用空间小的类型
