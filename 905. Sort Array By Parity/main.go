package main

import "fmt"

func main() {
	A := []int{3, 1, 2, 4}

	// var oddArr, evenArr []int
	// for _, v := range A {
	// 	if v%2 == 0 {
	// 		evenArr = append(evenArr, v)
	// 	} else {
	// 		oddArr = append(oddArr, v)
	// 	}
	// }
	// //合并slice的用法
	// fmt.Println(append(evenArr, oddArr...))

	//第二种解法
	j := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			A[i], A[j] = A[j], A[i]
			j++
		}
	}

	fmt.Println(A)

}
