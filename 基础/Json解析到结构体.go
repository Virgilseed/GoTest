package 基础

import (
	"fmt"
	"os"
)

type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

func main() {

	fmt.Println("Are you ok?")
	os.Stdout.WriteString("yes\n")
	var a int
	fmt.Println("请输入a：")
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
