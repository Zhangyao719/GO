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