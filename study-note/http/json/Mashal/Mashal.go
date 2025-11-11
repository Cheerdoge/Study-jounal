package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Type    string
	City    string `json:"Address"`
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	fmt.Println(vc)
	//序列化
	s, err := json.MarshalIndent(vc, "", "")
	if err != nil {
		fmt.Println("序列化失败：", err)
		return
	}
	fmt.Println(string(s))
	//已知数据的反序列化
	var x VCard
	err1 := json.Unmarshal(s, &x)
	if err1 != nil {
		fmt.Println("反序列化失败：", err)
	}
	fmt.Println(x)
	//未知数据的反序列化
	var f interface{}
	err2 := json.Unmarshal(s, &f)
	if err2 != nil {
		fmt.Println("反序列化失败：", err)
	}
	m := f.(map[string]interface{})
	fmt.Println(m)
}
