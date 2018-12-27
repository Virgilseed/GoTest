package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	var start, end int
	fmt.Println("请输入起始页（ >= 1）:")
	fmt.Scan(&start)
	fmt.Println("请输入结束页( >= 起始页):")
	fmt.Scan(&end)
	DoWork(start, end)
}
func DoWork(start int, end int) {
	fmt.Printf("准备爬第 %d 页到第 %d 页网址:\n", start, end)

	for i := start; i <= end; i++ {
		SpiderHomePage(i)
	}
}
func SpiderHomePage(i int) {
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("准备爬第 %d 页的网址: %s\n", i, url)

	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet error = ", err)
		return
	}

	re := regexp.MustCompile("<h1 class=\"dp-b\"><a href=\"(?s:(.*?))\"")
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	joyUrls := re.FindAllStringSubmatch(result, -1)
	fmt.Println("joyUrls", joyUrls)

	for _, data := range joyUrls {
		fmt.Println("url = ", data)
	}

}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}

	defer resp.Body.Close()
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf)
	}
	return
}
