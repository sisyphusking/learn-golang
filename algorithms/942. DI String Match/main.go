package main

import "fmt"

func main() {
	s := "IDID"
	list := diStringMatch(s)
	fmt.Println(list)
}

func diStringMatch(S string) []int {

	i, j := 0, len(S)
	ret := make([]int, len(S)+1)
	for k := 0; k < len(S); k++ {
		if S[k] == 'I' {
			ret[k] = i
			i++
		} else {
			ret[k] = j
			j--
		}
	}
	ret[len(S)] = i
	return ret

}
