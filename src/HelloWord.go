package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("Hello go")
	//
	//sum := 1
	//for i := 0; i < 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
	//for ; sum < 100;{
	//	sum += sum
	//}
	//fmt.Println(sum)

	//for sum < 100 {
	//	sum += sum
	//}
	//fmt.Println(sum)
	//
	//fmt.Println(sqrt(2), sqrt(-4))
	//
	//fmt.Println(
	//	pow(3, 2, 10),
	//	pow(3, 3, 20),
	//)

	//switch os := runtime.GOOS; os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux.")
	//default:
	//	fmt.Printf("%s.\n", os)
	//}
	//
	//fmt.Println("When's Saturday?")
	//today := time.Now().Weekday()
	//switch time.Saturday {
	//case today + 0:
	//	fmt.Println("Today.")
	//case today + 1:
	//	fmt.Println("Tomorrow.")
	//case today + 2:
	//	fmt.Println("In two days.")
	//default:
	//	fmt.Println("Too far away.")
	//}

	//defer fmt.Println("world")
	//fmt.Println("hello")

	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//fmt.Println("done")

	//i, j := 42, 2701
	//p := &i
	//fmt.Println(*p)
	//*p = 21
	//fmt.Println(i)
	//p = &j
	//*p = *p / 37
	//fmt.Println(j)

	//fmt.Println(Vertex{1, 2})
	//v := Vertex{1, 2}
	//v.X = 4
	//fmt.Println(v.X)

	//v := Vertex{1, 2}
	//p := &v
	//p.X = 10
	//fmt.Println(v)
	//fmt.Println(v1, *p, v2, v3)
	//
	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Print(a[0], a[1])
	//fmt.Print(a)
	//
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(primes)

	//var s []int = primes[1:4]
	//fmt.Println(s)
	//
	//names := [4]string{
	//	"John",
	//	"Paul",
	//	"George",
	//	"Ringo",
	//}
	//fmt.Println(names)

	//c := names[0:2]
	//b := names[1:3]
	//fmt.Println(c, b)
	//
	//b[0] = "XXX"
	//fmt.Println(c, b)
	//fmt.Println(names)

	//q := [7]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(q)
	//
	//r := []bool{true, false, true, true, false, true}
	//fmt.Println(r)

	//ss := []struct {
	//	i int
	//	b bool
	//}{
	//	{2, true},
	//	{3, false},
	//	{5, true},
	//	{7, true},
	//	{11, false},
	//	{13, true},
	//}
	//fmt.Println(ss)

	//sss := []int{2, 3, 5, 7, 11, 13}
	//printSlice(sss)
	//
	//// Slice the slice to give it zero length.
	//sss = sss[:0]
	//printSlice(sss)
	//
	//// Extend its length.
	//sss = sss[:7]
	//printSlice(sss)
	//
	//// Drop its first two values.
	//sss = sss[2:]
	//printSlice(sss)

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

type Vertex struct {
	X int
	Y int
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
