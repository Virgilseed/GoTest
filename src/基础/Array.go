package 基础

import (
	"fmt"
	"math"
)

func main() {
	//var s []int
	//printSlice(s)
	//
	//// append works on nil slices.
	//s = append(s, 0)
	//printSlice(s)
	//
	//// The slice grows as needed.
	//s = append(s, 1)
	//printSlice(s)
	//
	//// We can add more than one element at a time.
	//s = append(s, 2, 3, 4)
	//printSlice(s)
	//
	//for i, v := range pows{
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}
	//
	//pow := make([]int, 10)
	//for i := range pow {
	//	pow[i] = 1 << uint(i) // == 2**i
	//}
	//for _, value := range pow {
	//	fmt.Printf("%d\n", value)
	//}

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	fmt.Println(compute(math.Pow))

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

var pows = []int{1, 2, 4, 8, 16, 32, 64, 128}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
