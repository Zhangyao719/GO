package main

import "fmt"

func main() {
	var c1, c2 byte = 'a', '0'

	fmt.Println(c1, c2) // 97 48
	fmt.Printf("%c %c \n", c1, c2) // 'a', '0'
	
	// var c3 = '北' // error overflow
	var c3 int = '北'
	fmt.Printf("%c %d \n", c3, c3) // 北 21271
	
	var c4 byte = 10
	c5 := c4 + 'a'
	fmt.Printf("%c %d", c5, c5) // k 107
}

// * 1. 当直接输出byte时，就是输出了对应字符的ASC十进制码值
// * 2. 可以使用格式化输出 %c，转换成字符
// * 3. 如果字符对应的码值大于255，考虑使用int保存
// * 4. 字符可以进行运算