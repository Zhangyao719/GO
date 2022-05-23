# Go

## Hello world

```go
// * 每个文件必须归属于一个package
package main

// 内置提供格式化、输入、输出的函数
import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## 编译

- build: , 执行`go build xxx.go`， 会生成 xxx.exe二进制文件。

  ```go
  go build xxx.go
  
  // 参数选项：
  -o 指定生成的文件名
  ```

- 运行

  ```go
  ./xxx.exe
  ```

- 说明: 一个可直接运行的 Golang 程序必须有`main`包和入口`main`函数:

  ```bash
  # 没有main包报错:
  go run: cannot run non-main package
  
  # 没有main函数报错:
  # command-line-arguments
  runtime.main_main·f: function main is undeclared in the main package
  
  ```

- `go run`  命令也可以直接 编译并执行。

- `gofmt -w xxx.go` 格式化

## 转义字符

1. \t :  制表符 
2. \n :  换行
3. \\\ :  一个\
4. \\" : 一个"
5. \r : 一个回车，不是换行，会将后面的内容重新写在行首进行覆盖。

## DOS操作指令

> DOS: Disk Operating System 磁盘操作系统

- **目录操作**

1. `dir`: 查看目录
2. `md`: 创建文件夹 （多个目录 依次往后写）
3. `rd `: 删除空目录（`/q/s` 无需询问，删除目录下所有文件，`/s` 删除时 出现是否删除提示）
4. `cd /d f:` : 切换到F盘
5. `cd d:\demo` : 切换到D盘demo目录下
6. `cd \`: 返回根目录 

- **文件操作**

1. `echo`：输出字符串

   ```powershell
   # 将 "hello" 写入 D盘下的demo/test01/abc.txt文本
   echo hello > d:\demo\test01\abc.txt
   ```

2. `copy`：复制文件

   ```powershell
   # 将 abc.txt 文件复制到 test02 目录下
   copy abc.txt d:\demo\test02
   copy abc.txt d:\demo\test02\重命名.txt
   ```

3. `move`：移动文件（剪切）

   ```powershell、
   # 将 abc.txt 文件 剪切到 D盘/demo文件夹下
   move abc.txt d:\demo\
   ```

4. `del`： 删除文件

   ```powershell
   del abc.txt
   del *.txt # 删除所有 .txt 文件
   ```

5. `exit`：退出cmd

## 变量

### 变量使用三种方式

- 指定变量类型，若声明后不赋值，将使用默认值。

  ```go
  var num int = 10
  
  // 不赋值 int类型 默认0
  var num int
  ```

- 类型推导（根据值自行判断类型）

  ```go
  var num = 10.11
  ```

- 省略 var 注意：`:= ` 左侧的变量不应该是已经声明过的，否则编译失败。

  ```go
  name := 'Tom'
  ```

### 全局变量

```go
// 全局变量中 无法使用 :=

// 批量声明方式一：
var global1, global2, global3 = 1, 2, 3

// 批量声明方式二：
var (
	global1 = 1
    global2 = 2
    global2 = 3
)
```

## 数据类型

**基本数据类型**

数值型: 整数（int / int8 / int16 / int32 / int64 / uint / uint8 / uint16 / uint32 / uint64 / byte ）

字符型：没有专门的字符型，使用byte来保存单个字母字符

布尔型

字符串

**派生/复杂数据类型**

指针(Pointer)、数组、结构体(struct)、管道(Channel)、函数、切片(slice)、接口(interface)、map

### 整数

| 类型   | 有无符号（区分正负） | 占用存储空间                         | 表数范围                           | 备注                     |
| ------ | -------------------- | ------------------------------------ | ---------------------------------- | ------------------------ |
| int8   | 有                   | 1字节                                | -128 ~ 127                         |                          |
| int16  | 有                   | 2字节                                | -2^15 ~ 2^15 -1                    |                          |
| int32  | 有                   | 4字节                                | -2^31 ~ 2^31 -1                    |                          |
| int64  | 有                   | 8字节                                | -2^63 ~ 2^63 -1                    |                          |
| uint8  | 无                   | 1字节                                | 0 ~ 255                            |                          |
| uint16 | 无                   | 2字节                                | 0 ~ 2^16 - 1                       |                          |
| uint32 | 无                   | 4字节                                | 0 ~ 2^32 - 1                       |                          |
| uint64 | 无                   | 8字节                                | 0 ~ 2^64 - 1                       |                          |
| int    | 有                   | 32位系统 4个字节<br>64位系统8个字节  | -2^31 ~ 2^31 -1<br>-2^63 ~ 2^63 -1 |                          |
| uint   | 无                   | 32位系统 4个字节<br/>64位系统8个字节 | 0 ~ 2^32 - 1<br>0 ~ 2^64 - 1       |                          |
| rune   | 有                   | 与 int32 一样                        | 0 ~ 2^32 - 1                       | 表示一个Unicode码        |
| byte   | 无                   | 与 uint8 一样                        | 0 ~ 255                            | 当要存储字符时，选用byte |

**查看变量占用字节大小和数据类型**

`unsafe.Sizeof(num)`: 获取占用字节数

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int64 = 12 
	fmt.Printf("num 类型 %T 字节数 %d", num, unsafe.Sizeof(num))
}

// num → %T，unsafe.Sizeof(num) → %d
// output: num 类型 int64 字节数 8
```

### 浮点

| 类型           | 占用存储空间 | 表数范围               |
| -------------- | ------------ | ---------------------- |
| 单精度 float32 | 4字节        | -3.403E38 ~ 3.403E38   |
| 双精度 float64 | 8字节        | -1.798E308 ~ 1.798E308 |

1. 浮点类型默认声明为 float64 类型；
2. 尾数部分可能丢失，造成精度丢失，首选float64；
3. 支持科学计数法。

### 字符（字母）

字符串就是一串固定长度的字符连接起来的字符序列。GO的字符串是由单个字节连接起来的。也就是说对于传统的字符串是由字符组成的，而**GO的字符串不同，它是由字节组成的，本质就是一个整数。**

GO没有专门的字符类型，**使用 UTF-8 编码**，如果要存储单个字符（字母），一般使用`byte`来保存。

1. 字符使用单引号''表示；
2. 当直接输出byte时，就是**输出了对应字符的 ASCⅡ 十进制码值**；
3. 可以使用格式化输出 %c，转换成字符；
4. 如果字符对应的码值大于255，考虑使用uint保存；
5. 字符是可以运算的，相当于一个整数，因为它都对应有Unicode码。

```go
package main
import "fmt"
func main() {
	var c1, c2 byte = 'a', '0'

	fmt.Println(c1, c2) // 97 48
	fmt.Printf("%c %c \n", c1, c2) // 'a', '0'
	
	// var c3 = '北' // error overflow
	var c3 int = '北'
	fmt.Printf("%c %d", c3, c3) // 北 21271
}
```

```go
// 字符运算：
func main()  {
    var c4 byte = 10
	c5 := c4 + 'a'
	fmt.Printf("%c %d", c5, c5) // k 107
}
```



#### [关于字符的介绍](https://www.ruanyifeng.com/blog/2007/10/ascii_unicode_and_utf-8.html)

- [ASCⅡ](http://c.biancheng.net/c/ascii/) 共128个字符（第一位统一是0）

- **Unicode**

  对 ASCII 的补充，但是它只规定了符号的二进制代码，却没有规定这个二进制代码应该如何存储。

  比如一个汉字可能需要3个字节，但是计算机无法知道这段字符表示一个汉字符号，还是分别表示三个其他符号。第二个问题：会带来字节的浪费，无故扩大文件的存储大小。

- **UTF-8**

  **UTF-8 是 Unicode 的实现方式之一**。是针对Unicode的一种**可变长度**字符编码，**根据不同的符号而变化字节长度。**可以用来表示Unicode标准中的任何字符，而且其编码中的第一个字节仍与[ASCII](https://baike.baidu.com/item/ASCII/309296)相容，使得原来处理ASCII字符的软件无须或只进行少部分修改后，便可继续使用。

### 字符串

GO的字符串是由单个字节连接起来的（使用UTF-8编码标识的Unicode文本）。

GO中的**字符串可以重新赋值，但不能修改**。

 ```go
func main() {
	str := "字符串是不可变的"
	str[0] = "呵呵" // error cannot assign to str[0] (value of type byte)
	str = "哈哈" // success

	fmt.Println(str)
}
 ```

- 使用双引号"" : 有特殊字符 需要转义 
- 使用反引号`` : 内容全部视为字符串

```go
func main() {
	str1 := "我叫\"王多鱼\""
    str2 := `
    	庄强:"王多鱼是我哥们儿"
    	王多鱼:"请你一定要加入我的团队 大聪明~！"
    `
	fmt.Println(str)
}
```

- 超长字符串换行拼接（+号 在结尾）

```go
func main() {
	str3 := "哈哈" + "呵呵"
	str3 += "嘿嘿"
    // 字符串的拼接(超长字符串注意 +号 在结尾)
	str4 := "hello" + "world" + "hello" + "world" + "hello" + "world" +
	"hello" + "world" + "hello" + "world" +
	"hello" + "world"

	fmt.Println(str3, str4)
}
```

### 基本数据类型的默认值

整数型：0

浮点型：0

字符串：""

布尔值：false

### 基本数据转换

- **数字转换**

  高精度在向低精度转换时，如果高精度超出低精度范围，编译并不会报错，而会做溢出处理。

  ```go
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
  	fmt.Println("n5 \n", n5) // 231
  	
  	var n6 int32 = 300
  	var n7 int8 = int8(n6) + 127 // success 编译通过 溢出处理
  	// var n8 int8 = int8(n6) + 128 // error 128本身已超出 int8 范围
  	fmt.Println("n7", n7) // -85
  }
  ```

- **转字符串**