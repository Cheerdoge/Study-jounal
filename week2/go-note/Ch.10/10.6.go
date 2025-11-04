// 定义结构体 employee，它有一个 salary 字段，给这个结构体定义一个方法 giveRaise 来按照指定的百分比增加薪水
package main

import "fmt"

type employee struct {
	name   string
	salary int
}

func (p *employee) giveRaise() int {
	p.salary = p.salary * 2
	return p.salary
}

func main() {
	employee1 := new(employee)
	employee1.name = "cheerdoge"
	employee1.salary = 1000
	fmt.Printf("hi，%s,你的工资原本是：%d\n", employee1.name, employee1.salary)
	fmt.Println("但是你潜入公司系统给自己加薪了!")
	employee1.giveRaise()
	fmt.Println("现在你的薪水是：", employee1.salary)
}
