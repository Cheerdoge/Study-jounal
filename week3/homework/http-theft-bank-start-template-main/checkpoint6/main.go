package main

import (
	"fmt"

	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}

	req.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiQ2hlZXJkb2dlIiwiaWF0IjoxNzYyODY4NjE2LCJuYmYiOjE3NjI4Njg2MTZ9.siKywu0RdUEwq5vwB33ijTkeO3B3Rj5bm7ZfIZx5Y8E")

	resp, err := req.SendRequest()
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.ShowHeader()
	resp.ShowBody()
}
