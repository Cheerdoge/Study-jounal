package main

import (
	"fmt"

	"github.com/weiji6/hacker-support/encrypt"
	"github.com/weiji6/hacker-support/httptool"
)

func main() {
	data, err := encrypt.Base64Decode("c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ==")
	if err != nil {
		fmt.Println("解析失败：", err)
		return
	}
	fmt.Println(data)

	error_code := []byte("for {go func(){time.Sleep(1*time.Hour)}()}")
	secret_key := []byte("MuxiStudio203304")
	code, err := encrypt.AESEncryptOutInBase64(error_code, secret_key)
	if err != nil {
		fmt.Println("加密失败：", err)
		return
	}

	req, err := httptool.NewRequest(
		httptool.PUTMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate",
		string(code),
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}
	req.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoiQ2hlZXJkb2dlIiwiaWF0IjoxNzYyODY4NjE2LCJuYmYiOjE3NjI4Njg2MTZ9.siKywu0RdUEwq5vwB33ijTkeO3B3Rj5bm7ZfIZx5Y8E")

	// write your code below
	resp, err := req.SendRequest()
	if err != nil {
		fmt.Println(err)
		return
	}

	resp.ShowHeader()
	resp.ShowBody()
}
