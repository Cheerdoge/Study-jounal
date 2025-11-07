// 问ai提到的切片扩容是1024前超出内存则乘2并拷贝到一个新的地址，
// map则是通过哈希表实现的，扩容时会创建一个更大的哈希表，并将原有的键值对重新哈希到新的表中，每一次操作移动一点点
// 但是the way to go 中提到map扩容是多一增一
// 所以我选择实践出真知
package main

import "fmt"

func main() {
	s := make([]int, 1)
	s[0] = 1
	for i := 0; i < 10; i++ {
		cap0 := cap(s)
		s = append(s, i)
		if cap0 != cap(s) {
			fmt.Println("cap(s)=", cap(s))
		}

	}

}
