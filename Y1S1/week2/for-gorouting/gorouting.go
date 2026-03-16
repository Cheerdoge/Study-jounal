// 真没招了，我真不知道不用额外空间怎么给结果排序
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type out struct {
	id  int
	num int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan out)
	for i := 1; i < 21; i++ {
		go func() {
			sleeptime := rand.Intn(1000)
			time.Sleep(time.Duration(sleeptime) * time.Millisecond)
			num := rand.Intn(100)
			result := out{i, num}
			ch <- result
		}()
	}

	slice := make([]out, 0, 20)

	for i := 0; i < 20; i++ {
		result := <-ch
		fmt.Printf("gorouting id: %d, random num: %d\n", result.id, result.num)
		slice = append(slice, result)
	}

	fmt.Printf("\n")

	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j].id > slice[j+1].id {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	for _, value := range slice {
		fmt.Printf("gorouting id: %d, random num: %d\n", value.id, value.num)
	}
}
