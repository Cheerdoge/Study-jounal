// 使用递归函数从 10 打印到 1。
package main

import "fmt"

func main() {
	dofunc(10)
}

func dofunc(n int) {
	fmt.Println(n)
	if n > 1 {
		dofunc(n - 1)
	}
}
