package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

func main() {
	//v := Vertex{3, 4}
	//v.Scale(10)
	//fmt.Println(v.Abs())

	//var i I
	//describe(i)
	//i.M()

	//var i2 interface{}
	//describes(i2)
	//i2 = 42
	//describes(i2)
	//
	//i2 = "hello"
	//describes(i2)

	//var i interface{} = "hello"
	//
	//s := i.(string)
	//fmt.Println(s)
	//
	//s, ok := i.(string)
	//fmt.Println(s, ok)
	//
	//f, ok := i.(float64)
	//fmt.Println(f, ok)
	//
	//f = i.(float64) // panic
	//fmt.Println(f)

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

}

func describes(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
