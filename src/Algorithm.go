package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 插入排序
// charu(buf)
//希尔排序
// xier(buf)
//快速排序
// kuaisu(buf)
// 归并排序
// guibing(buf)
// 堆排序

/**
 * 冒泡排序
 */
func maopao1(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		flag := false
		for j := 1; j < len(buf)-i; j++ {
			if buf[j-1] > buf[j] {
				times++
				buf[j-1], buf[j] = buf[j], buf[j-1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println("maopao times: ", times)
	fmt.Println(buf)
}

func maopao(arrray []int) {
	changeTimes := 0
	compareTimes := 0
	for i := 0; i < len(arrray)-1; i++ {
		for j := 0; j < len(arrray)-i-1; j++ {
			compareTimes++
			if arrray[j] > arrray[j+1] {
				changeTimes++
				arrray[j], arrray[j+1] = arrray[j+1], arrray[j]
			}
		}
	}
	fmt.Println("冒泡比较次数: ", compareTimes)
	fmt.Println("冒泡交换次数: ", changeTimes)
	fmt.Println(arrray)
}

// 选择排序
func xuanze(arr []int) {
	changeTimes := 0
	compareTimes := 0
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
			compareTimes++
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
			changeTimes++
		}
	}
	fmt.Println("选择比较次数: ", compareTimes)
	fmt.Println("选择交换次数: ", changeTimes)
	fmt.Println(arr)
}

// 插入排序
func charu(buf []int) {
	changeTimes := 0
	compareTimes := 0
	for i := 1; i < len(buf); i++ {
		for j := i; j > 0; j-- {
			compareTimes++
			if buf[j] < buf[j-1] {
				changeTimes++
				buf[j-1], buf[j] = buf[j], buf[j-1]
			} else {
				break
			}
		}
	}
	fmt.Println("插入比较次数: ", compareTimes)
	fmt.Println("插入交换次数: ", changeTimes)
	fmt.Println(buf)
}

// 希尔排序
func xier(buf []int) {
	times := 0
	tmp := 0
	length := len(buf)
	incre := length
	// fmt.Println("buf: ", buf)
	for {
		incre /= 2
		for k := 0; k < incre; k++ { //根据增量分为若干子序列
			for i := k + incre; i < length; i += incre {
				for j := i; j > k; j -= incre {
					// fmt.Println("j: ", j, " data: ", buf[j], " j-incre: ", j-incre, " data: ", buf[j-incre])
					times++
					if buf[j] < buf[j-incre] {
						tmp = buf[j-incre]
						buf[j-incre] = buf[j]
						buf[j] = tmp
					} else {
						break
					}
				}
				// fmt.Println("middle: ", buf)
			}
			// fmt.Println("outer: ", buf)
		}
		// fmt.Println("outer outer: ", buf, " incre: ", incre)

		if incre == 1 {
			break
		}
	}
	// fmt.Println("after: ", buf)
	fmt.Println("xier times: ", times)
}

// 快速排序
func kuaisu(buf []int) {
	kuai(buf, 0, len(buf)-1)
}

func kuai(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j, key := l, r, a[l] //选择第一个数为key
	for i < j {
		for i < j && a[j] > key { //从右向左找第一个小于key的值
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}

		for i < j && a[i] < key { //从左向右找第一个大于key的值
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	//i == j
	a[i] = key
	kuai(a, l, i-1)
	kuai(a, i+1, r)
}

//归并排序
func guibing(buf []int) {
	tmp := make([]int, len(buf))
	merge_sort(buf, 0, len(buf)-1, tmp)
}

func merge_sort(a []int, first, last int, tmp []int) {
	if first < last {
		middle := (first + last) / 2
		merge_sort(a, first, middle, tmp)       //左半部分排好序
		merge_sort(a, middle+1, last, tmp)      //右半部分排好序
		mergeArray(a, first, middle, last, tmp) //合并左右部分
	}
}

func mergeArray(a []int, first, middle, end int, tmp []int) {
	// fmt.Printf("mergeArray a: %v, first: %v, middle: %v, end: %v, tmp: %v\n",
	//     a, first, middle, end, tmp)
	i, m, j, n, k := first, middle, middle+1, end, 0
	for i <= m && j <= n {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}
	for i <= m {
		tmp[k] = a[i]
		k++
		i++
	}
	for j <= n {
		tmp[k] = a[j]
		k++
		j++
	}

	for ii := 0; ii < k; ii++ {
		a[first+ii] = tmp[ii]
	}
	// fmt.Printf("sort: buf: %v\n", a)
}

// 堆排序
func duipai(buf []int) {
	temp, n := 0, len(buf)

	for i := (n - 1) / 2; i >= 0; i-- {
		MinHeapFixdown(buf, i, n)
	}

	for i := n - 1; i > 0; i-- {
		temp = buf[0]
		buf[0] = buf[i]
		buf[i] = temp
		MinHeapFixdown(buf, 0, i)
	}
}

func MinHeapFixdown(a []int, i, n int) {
	j, temp := 2*i+1, 0
	for j < n {
		if j+1 < n && a[j+1] < a[j] {
			j++
		}

		if a[i] <= a[j] {
			break
		}

		temp = a[i]
		a[i] = a[j]
		a[j] = temp

		i = j
		j = 2*i + 1
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//array := make([]int, 20)
	array := []int{878, 412, 304, 220, 828, 142, 68, 625, 778, 462, 231, 679, 478, 592, 935, 492, 698, 414, 206, 389}
	//for i := 0; i < 20; i++ {
	//	array[i] = rand.Intn(1000)
	//}

	fmt.Println("排序前：")
	fmt.Println(array)
	fmt.Println("排序后：")
	//maopao1(array)
	//maopao(array)
	xuanze(array)
}