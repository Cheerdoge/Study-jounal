package main

import "fmt"

type Request struct {
	name   string
	days   int
	date   string
	reason string
}

type Handler interface {
	Handle(Request)
	SetNext(Handler)
	ReturnLevel() int
}

//1~3
type Student struct {
	Name        string
	NextHandler Handler
	Level       int
}

//4~6
type Teacher struct {
	Name        string
	NextHandler Handler
	Level       int
}

//7~9
type BigHand struct {
	Name        string
	NextHandler Handler
	Level       int
}

func (s *Student) SetNext(handler Handler) {
	s.NextHandler = handler
	return
}

func (s *Student) Handle(request Request) {
	fmt.Println("学生" + s.Name + "开始审批" + request.name + "的请假请求")
	if s.Level >= request.days {
		fmt.Println("学生" + s.Name + "批准了" + request.name + "的请假请求")
	} else {
		fmt.Println("学生" + s.Name + "无法批准" + request.name + "的请假请求，已转交")
		s.NextHandler.Handle(request)
	}
}

func (s *Student) ReturnLevel() int {
	return s.Level
}

func (t *Teacher) SetNext(handler Handler) {
	t.NextHandler = handler
	return
}

func (t *Teacher) Handle(request Request) {
	fmt.Println("老师" + t.Name + "开始审批" + request.name + "的请假请求")
	if t.Level >= request.days {
		fmt.Println("老师" + t.Name + "批准了" + request.name + "的请假请求")
	} else {
		fmt.Println("老师" + t.Name + "无法批准" + request.name + "的请假请求，已转交")
		t.NextHandler.Handle(request)
	}
}

func (t *Teacher) ReturnLevel() int {
	return t.Level
}

func (b *BigHand) SetNext(handler Handler) {
	b.NextHandler = handler
	return
}

func (b *BigHand) Handle(request Request) {
	fmt.Println("大手子" + b.Name + "开始审批" + request.name + "的请假请求")
	if b.Level >= request.days {
		fmt.Println("大手子" + b.Name + "批准了" + request.name + "的请假请求，速速感谢大手子")
	} else {
		fmt.Println("大手子" + b.Name + "无法批准" + request.name + "的请假请求，回家吧孩子")
	}
}

func (b *BigHand) ReturnLevel() int {
	return b.Level
}

func NewStudent(name string, level int) *Student {
	return &Student{Name: name, Level: level}
}

func NewTeacher(name string, level int) *Teacher {
	return &Teacher{Name: name, Level: level}
}

func NewBigHand(name string, level int) *BigHand {
	return &BigHand{Name: name, Level: level}
}

func NewList() []Handler {
	return []Handler{}
}

func AddHandler(list []Handler, handler Handler) []Handler {
	if len(list) == 0 || list[0].ReturnLevel() > handler.ReturnLevel() {
		list = append([]Handler{handler}, list...)
		if len(list) > 1 {
			list[0].SetNext(list[1])
		}
		return list
	}
	for i := 0; i < len(list)-1; i++ {
		if list[i].ReturnLevel() < handler.ReturnLevel() {
			list = append(list[:i+1], append([]Handler{handler}, list[i+1:]...)...)
			list[i].SetNext(list[i+1])
			list[i+1].SetNext(list[i+2])
			return list
		}
	}
	list = append(list, handler)
	list[len(list)-2].SetNext(list[len(list)-1])
	return list
}

func main() {
	student := NewStudent("人民的好班长", 3)
	teacher := NewTeacher("伟大的辅导员", 6)
	bighand := NewBigHand("？？？", 9)

	handlerList := NewList()
	handlerList = AddHandler(handlerList, teacher)
	handlerList = AddHandler(handlerList, bighand)
	handlerList = AddHandler(handlerList, student)

	request := Request{
		name:   "小明",
		days:   7,
		date:   "2024-12-01 to 2024-12-07",
		reason: "回家探亲",
	}
	handlerList[0].Handle(request)
}
