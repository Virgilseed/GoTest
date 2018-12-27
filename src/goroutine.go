package main

import (
	"fmt"
	"time"
)

func NewTask() {
	for {
		fmt.Println("this is NewTask")
		time.Sleep(time.Second)
	}
}

func main() {
	go NewTask()
	for {
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second)
	}

}
