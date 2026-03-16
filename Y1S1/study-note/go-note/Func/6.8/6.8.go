// 在 main 函数中写一个用于打印 Hello World 字符串的匿名函数并赋值给变量 fv，然后调用该函数并打印变量fv 的类型。
package main

import "fmt"

func main() {
	g := func() {
		fmt.Println("Hello World")
	}
	g()
}
