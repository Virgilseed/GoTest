package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer file.Close()
	//var buf string
	//for i := 0; i < 10; i++ {
	//	buf = fmt.Sprintf("i = %d\n", i)
	//	_, err := file.WriteString(buf)
	//	if err != nil {
	//		fmt.Println("err = ", err)
	//	}
	//}
	bu := bufio.NewReader(file)
	for {
		buf, err1 := bu.ReadBytes('\n')
		if err1 != nil {
			if err1 == io.EOF {
				break
			}
			fmt.Println("err1 = ", err1)
		}
		fmt.Printf("buf = #%s\n", string(buf))
	}

}

func ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer file.Close()
	buf := make([]byte, 1024*2)
	n, err1 := file.Read(buf)
	if err1 != nil && err1 != io.EOF {
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = ", string(buf))
	fmt.Println("n = ", n)

}

func main() {
	path := "./文件.txt"
	WriteFile(path)
	ReadFile(path)
}
