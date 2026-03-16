// 删减成计算到12的阶乘，因为big包还不了解
// 啊啊递归好像忘干净了
package main

import "fmt"

func dofunc(n int) (result int) {
	if n == 1 {
		result = 1
	} else {
		result = n * dofunc(n-1)
	}
	return
}

func main() {
	for i := 1; i <= 12; i++ {
		result := dofunc(i)
		fmt.Printf("%d的阶乘结果是：%d\n", i, result)
	}
}

//这里原本i:=0的，但是没有添加i==0的判断条件，导致出错
