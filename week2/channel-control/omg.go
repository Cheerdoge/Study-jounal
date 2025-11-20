// 求助ai知道了要用信号流
// 怎么控制谁先真不会了
package main

import (
	"fmt"
	"time"
)

func main() {
	lingpai1 := make(chan struct{}, 1)
	lingpai2 := make(chan struct{}, 1)

	lingpai1 <- struct{}{}

	go func() {
		s := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		for i := 0; i < 24; i++ {
			<-lingpai1
			fmt.Printf("%c", s[i])
			i++
			fmt.Printf("%c", s[i])
			lingpai2 <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < 26; i++ {
			<-lingpai2
			fmt.Printf("%d", i)
			i++
			fmt.Printf("%d", i)
			lingpai1 <- struct{}{}
		}
	}()

	time.Sleep(10 * time.Second)

}
