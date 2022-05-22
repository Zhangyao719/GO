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

## 