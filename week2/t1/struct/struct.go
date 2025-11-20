package main

import "fmt"

type text struct {
	num int
	s   bool
}

func main() {
	a := text{10, true}
	b := text{10, true}
	c := text{9, true}
	fmt.Println(a == b)
	fmt.Println(a == c)

}
