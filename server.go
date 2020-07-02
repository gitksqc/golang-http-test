package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", testgo)
	
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func testgo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, " connected success.")
	fmt.Println("url: ", r.URL.Path)
	fmt.Println("method: ", r.Method)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)
	
	resp_str := "hello world"
	fmt.Println("length of response: ", len(resp_str))
	w.Write([]byte(resp_str))
}