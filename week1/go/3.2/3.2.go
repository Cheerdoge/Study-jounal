// 1.解释：长度看的是字符数,其中逗号应该是中文字符，加上世界，每个算三个字符，共14
// 2.解释：range就直接看这段字符串的字符数，共8个，输出8
package main

import "fmt"

func main() {
	var s string = "hello，世界"
	fmt.Println("len(s) = ", len(s))
	var num int = 0
	for range s {
		num++
	}
	fmt.Println("num = ", num)
}
