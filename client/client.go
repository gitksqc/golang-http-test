// client.go
package main

import (
	"io"
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Get("http://localhost:8080/go")
	if err != nil {
		fmt.Println("Connect failed...")
		return
	}
	defer resp.Body.Close()
	fmt.Println("resp status: ", resp.Status)
	fmt.Println("resp header: ", resp.Header.Get(""))
	
	buf := make([]byte, 2)
    
	for {
		n, err := resp.Body.Read(buf)	
		fmt.Println("err: ", err)
		if err == io.EOF {
			fmt.Println(string(buf[:n]))
			fmt.Println("the end of the response.")
			return
		} else if err != nil {
			break
		} else {
			fmt.Println(string(buf[:n]))
			continue
		}
		
	}
}
