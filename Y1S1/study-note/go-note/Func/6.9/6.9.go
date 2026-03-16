// 不使用递归但使用闭包改写第 6.6 节中的斐波那契数列程序
package main

import "fmt"

func dowork() func() (result int) {
	a := 0
	b := 1
	return func() (result int) {

		result = a
		a, b = b, a+b
		return
	}
}

func main() {
	f := dowork()
	for i := 0; i <= 10; i++ {
		fmt.Printf("第%d项的值为：%d\n", i, f())
	}
}

//写这个时把a,b写进闭包内了，导致a,b每次调用都被重新初始化
