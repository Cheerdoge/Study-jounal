// 其实没看太懂题目的意思，暂时理解成匿名函数运行几次输出几
package main

import "fmt"

func main() {
	f := dowork()
	for i := 0; i < 10; i++ {
		f()
	}
}

func dowork() func() {
	num := 0
	return func() {
		num++
		fmt.Printf("当前函数运行了%d次\n", num)
	}
}

//编写时习惯性写了return语句,但是没有返回值的函数不需要写
//差点让dowork函数的返回函数类型和匿名函数类型对不上
