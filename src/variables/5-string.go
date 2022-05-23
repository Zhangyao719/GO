package main

import "fmt"

func main() {
	// * 1. 字符串可以重新被赋值，但不能修改
	str := "字符串是不可变的"

	// str[0] = "呵呵" // error cannot assign to str[0] (value of type byte)

	str = "哈哈"

	// * 2. 双引号"" 和 反引号``的使用
	str1 := "我叫\"王多鱼\""
	str2 := `
		庄强:"王多鱼是我哥们儿"
		王多鱼:"请你一定要加入我的团队 大聪明~！"
	`

	// * 3. 字符串的拼接(超长字符串注意 +号 在结尾)
	str3 := "哈哈" + "呵呵"
	str3 += "嘿嘿"
	str4 := "hello" + "world" + "hello" + "world" + "hello" + "world" +
	"hello" + "world" + "hello" + "world" +
	"hello" + "world"

	fmt.Println(str, str1, str2, str3, str4)
}