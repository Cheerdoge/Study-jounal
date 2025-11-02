// 1. 解释：使用make创建一个切片后，若没有指定容量，那么他的容量就等于长度
// 2. 解释：切片的部分[low:high:max]  从下标为low开始，len=high-low  cap=max-low，如果没有指定max，那么容量就等于原来数组的容量减去low
package main

import "fmt"

func main() {
	s := make([]byte, 5)
	fmt.Println("len(s)=", len(s))
	fmt.Println("cap(s)=", cap(s))
	s = s[2:4]
	fmt.Println("len(s)=", len(s))
	fmt.Println("cap(s)=", cap(s))
}
