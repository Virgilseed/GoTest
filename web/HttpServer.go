package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)

}

func handle(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Header)
	fmt.Println(request.Method)
	fmt.Println(request.Body)
	fmt.Println(request.URL)
	writer.Write([]byte("hello go"))

}
