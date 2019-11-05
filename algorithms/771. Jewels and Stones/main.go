package main

import (
	"fmt"
	"strings"
)

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}

func numJewelsInStones(J string, S string) int {
	var count int
	for _, s := range S {
		if strings.Contains(J, string(s)) {
			count += 1
		}
	}
	return count
}
