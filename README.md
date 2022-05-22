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