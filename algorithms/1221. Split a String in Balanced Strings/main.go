package main

import "fmt"

func main() {
	s := "RLRRLLRLRL"
	fmt.Println(balancedStringSplit(s))
}

func balancedStringSplit(s string) int {
	var L, R, cnt int

	for i := 0; i < len(s); i++ {
		if L == R {
			cnt += 1
			L = 0
			R = 0
		}
		if s[i] == 'L' {
			L++
		} else {
			R++
		}
	}
	return cnt
}
