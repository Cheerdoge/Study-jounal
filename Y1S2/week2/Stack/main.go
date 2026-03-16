package main

import (
	anystack "Stack/any_stack"
	"fmt"
)

func main() {
	// 整数栈
	intStack := anystack.InitAnyStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack.Pop()) // 输出: 3 true

	// 字符串栈
	stringStack := anystack.InitAnyStack[string]()
	stringStack.Push("hello")
	stringStack.Push("world")

	fmt.Println(stringStack.Pop()) // 输出: world true
}
