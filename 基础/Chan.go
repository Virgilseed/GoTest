package 基础

import (
	"fmt"
	"time"
)

func main() {
	//a := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int)
	//go sum(a[:len(a)/2], c)
	//go sum(a[len(a)/2:], c)
	//x, y := <-c, <-c // 从 c 中获取
	//
	//fmt.Println(x, y, x+y)

	//go say("world")
	//say("hello")

	//c := make(chan int, 2)
	//c <- 1
	//c <- 2
	//c <- 3
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
