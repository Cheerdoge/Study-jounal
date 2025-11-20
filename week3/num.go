package main

import "fmt"

func permute(nums []int) [][]int {
	// insert your code
	result := make([][]int, 0)
	var dowork func(int)
	dowork = func(count int) {
		if count == len(nums) {
			value := make([]int, len(nums))
			copy(value, nums)
			result = append(result, value)
		}
		for i := count; i < len(nums); i++ {
			nums[count], nums[i] = nums[i], nums[count]
			dowork(count + 1)
			nums[count], nums[i] = nums[i], nums[count]
		}
	}
	dowork(0)
	return result
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	testSlice := make([]int, n)
	// 标准输入n个不重复的数字
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &testSlice[i])
	}
	res := permute(testSlice)
	fmt.Println(res)
}
