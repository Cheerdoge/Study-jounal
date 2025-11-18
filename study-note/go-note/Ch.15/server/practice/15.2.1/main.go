package main

import (
	"fmt"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In hello")
	name := req.URL.Path[len("/hello/"):]
	fmt.Fprint(w, "hello ", strings.ToUpper(name))
}

func shorthello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In shorthello")
	name := req.URL.Path[len("/shorthello/"):]
	fmt.Fprint(w, "hello ", strings.ToLower(name))
}

func main() {
	http.HandleFunc("/hello/", hello)
	http.HandleFunc("/shorthello/", shorthello)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
}
