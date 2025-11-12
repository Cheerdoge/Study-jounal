package main

import (
	"fmt"

	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/iris_sample",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiQ2hlZXJkb2dlIiwiaWF0IjoxNzYyODY4NjE2LCJuYmYiOjE3NjI4Njg2MTZ9.siKywu0RdUEwq5vwB33ijTkeO3B3Rj5bm7ZfIZx5Y8E")

	// write your code below
	resp, err := req.SendRequest()
	if err != nil {
		fmt.Println(err)
		return
	}

	err1 := resp.Save("D:\\code\\study-jounal\\week3\\homework\\http-theft-bank-start-template-main\\checkpoint5\\test.jpg")
	if err1 != nil {
		fmt.Println("下载失败:", err1)
		return
	}
	resp.ShowHeader()
	resp.ShowBody()
}
