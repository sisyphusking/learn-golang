package main

import "fmt"

func main() {
	a := []int{5, 1, 5, 2, 5, 3, 5, 4}
	m := make(map[int]int)
	r := 0
	for _, i := range a {
		_, ok := m[i]
		if ok {
			r = i
			break
		}
		m[i] = 0
	}

	fmt.Println(r)

}
