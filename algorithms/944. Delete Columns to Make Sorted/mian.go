package main

import "fmt"

func main() {
	a := []string{"cba", "daf", "ghi"}
	fmt.Println(minDeletionSize(a))
}

func minDeletionSize(A []string) int {

	m, n := len(A), len(A[0])
	//最小删除列
	res := 0
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if A[j][i] < A[j-1][i] {
				res++
				break
			}
		}
	}
	return res
}
