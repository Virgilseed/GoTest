package main

import (
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println("get error = ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status = ", resp.Status)
	fmt.Println("StatusCose = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)
	fmt.Println("Body = ", resp.Body)
	buff := make([]byte, 4*1024)
	var temp string
	for {
		n, err1 := resp.Body.Read(buff)
		if n == 0 {
			fmt.Println("read err1==", err1)
			break
		}

		temp += string(buff[:n])
	}
	fmt.Println("temp = ", temp)
}
