package main

import "fmt"

func permute(nums []int) [][]int {
	// insert your code

}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	testSlice := make([]int, n)
	// 标准输入n个不重复的数字
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &testSlice[i])
	}

	res := permute(testSlice)
	fmt.Println(res)
}
