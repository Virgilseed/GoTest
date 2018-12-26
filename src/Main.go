package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

//var c, python, java bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(1000))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	//fmt.Println(math.pi)
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))
	//var i int
	//var c, python, java = true, true, "no!"
	//fmt.Println(i, c ,python, java)
	//
	////var c, python, java = true, true, "no!"
	//fmt.Println(i, j, c ,python, java)

	//var i, j int = 1, 2
	//k := 3
	//c, python, java := true, false, false
	//fmt.Println(i, j, k, c, python, java)

	fmt.Printf("Type: %T Value:%v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value:%v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value:%v\n", z, z)

	//var i int
	//var f float64
	//var b bool
	//var s string
	//fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	v := 0.5i // 修改这里！
	fmt.Printf("v is of type %T\n", v)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
