package main

import (
	"fmt"
	"strings"
)

func main() {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	words := []string{"gin", "zen", "gig", "msg"}
	results := make(map[string]bool)
	for _, word := range words {
		// bword := []byte(word)
		var s strings.Builder
		for _, w := range word {
			s.WriteString(string(morseCodes[w-'a']))
		}
		// results = append(results, s.String())
		results[s.String()] = true
	}

	fmt.Println(len(results))
}
