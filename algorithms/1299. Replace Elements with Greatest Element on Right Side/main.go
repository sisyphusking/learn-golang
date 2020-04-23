package main

import "fmt"

//自己写的，性能不好
func replaceElements(arr []int) []int {
	var newArr []int
	for k, _ := range arr {

		if k == len(arr)-1 {
			break
		}
		rightArr := arr[k+1:]
		max := 1
		for _, i := range rightArr {
			if i > max {
				max = i
			}
		}
		newArr = append(newArr, max)
	}
	newArr = append(newArr, -1)
	return newArr
}

//耗时最低的，方式比较优雅，用max保存右边的最大值，逐个比对，超过max就替换，逻辑简单，主要是倒序，让人思维一下转不过来
func replaceElements1(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i -= 1 {
		el := arr[i]
		arr[i] = max
		if el > max {
			max = el
		}
	}
	return arr
}

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements1(arr))
}
