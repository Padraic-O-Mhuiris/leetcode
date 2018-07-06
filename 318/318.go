package main

import (
	"fmt"
)

func maxProduct(words []string) int {
	masks := make([]int, len(words))

	for i := range masks {
		fmt.Println(words[i])
		for j, b := range words[i] {
			fmt.Println(j, b)
			masks[i] |= 1 << uint(b-'a')
		}
	}
	maxProd := 0
	for i := range masks {
		for j := i + 1; j < len(masks); j++ {
			if masks[i]&masks[j] == 0 && (len(words[i])*len(words[j]) > maxProd) {

				maxProd = len(words[i]) * len(words[j])
			}
		}
	}
	return maxProd
}

func main() {

	s := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}

	fmt.Println(maxProduct(s))
}
