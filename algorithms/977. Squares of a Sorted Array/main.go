package main

import (
	"fmt"
	"sort"
)

func main() {

	sli := []int{-4, -1, 0, 3, 10}

	for index, ele := range sli {
		sli[index] = ele * ele
	}

	//实现接口来来调用Sort()方法
	// sort.Sort(ByValue(sli))

	//另一种方法
	sort.Ints(sli)
	fmt.Println(ByValue(sli))

}

//自定义类型，可以将sli赋值给它
type ByValue []int

func (b ByValue) Len() int {
	return len(b)
}

func (b ByValue) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b ByValue) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

//双指针
func SortArray(a []int) []int {
	var ret = make([]int, len(a))
	start, end := 0, len(a)-1
	index := len(a) - 1
	for start <= end {
		if a[start]*a[start] < a[end]*a[end] {
			ret[index] = a[end] * a[end]
			end--
		} else {
			ret[index] = a[start] * a[start]
			start++
		}
		index--
	}
	return ret
}
