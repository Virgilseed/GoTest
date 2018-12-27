package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var a [10]int
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
		fmt.Printf("%d, ", a[i])
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%d,", a[i])
	}
}
