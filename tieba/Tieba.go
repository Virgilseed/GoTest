package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var start, end int
	fmt.Printf("请输入起始页( >= 1):")
	fmt.Scan(&start)
	fmt.Printf("请输入结束页( >= 起始页):")
	fmt.Scan(&end)

	DoWork(start, end)

}

func DoWork(start int, end int) {
	fmt.Printf("正在爬取 %d 到 %d的页面\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPape(i, page)

	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			//fmt.Println("resp.Body.Read error = ", err2)
			return
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPape(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Printf("正在爬取第%d页网页：%s\n", i, url)
	result, err := HttpGet(url)
	if err != nil {
		//fmt.Println("HttpGet err = ", err)
		return
	}
	fileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create error = ", err1)
		return
	}
	f.WriteString(result)
	f.Close()
	page <- i
}
