package main

import "fmt"

func main() {
	a := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10}
	fmt.Println(uniqueOccurrences(a))
}

func uniqueOccurrences(A []int) bool {
	count := make(map[int]int, len(A))
	for _, a := range A {
		count[a]++
	}
	hasSeen := make(map[int]bool, len(count))
	for _, c := range count {
		//数组中有重复元素的话就立即退出
		if hasSeen[c] {
			return false
		}
		hasSeen[c] = true
	}
	//判断是否重复，方法二
	//temp := make(map[int]int)
	//for _, v := range count {
	//	temp[v]++
	//	if temp[v] > 1 { return false }
	//}
	return true
}

//最传统的方法
func uniqueOccurrences2(arr []int) bool {

	tmp := make(map[int]int)
	Arr := []int{}

	for _, v := range arr {
		tmp[v]++
	}
	for _, v := range tmp {
		Arr = append(Arr, v)
	}
	for i := 0; i < len(Arr); i++ {
		for j := i + 1; j < len(Arr); j++ {
			if Arr[i] == Arr[j] {
				return false
			}
		}
	}
	return true
}
