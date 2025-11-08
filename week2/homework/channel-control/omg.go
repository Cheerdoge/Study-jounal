// 求助ai知道了要用信号流
// 怎么控制谁先真不会了
package main

import (
	"fmt"
	"time"
)

func main() {
	lingpai := make(chan struct{}, 1)

	lingpai <- struct{}{}

	go func() {
		s := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
		for i := 0; i < 24; i++ {
			<-lingpai
			fmt.Printf("%c", s[i])
			i++
			fmt.Printf("%c", s[i])
			lingpai <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < 26; i++ {
			<-lingpai
			fmt.Printf("%d", i)
			i++
			fmt.Printf("%d", i)
			lingpai <- struct{}{}
		}
	}()

	time.Sleep(10 * time.Second)

}
