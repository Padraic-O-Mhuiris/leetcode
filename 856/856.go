package main

import (
	"fmt"
)

func scoreOfParentheses(S string) int {
	res, layers := 0, 0

	for i := range S {
		if S[i] == '(' {
			layers++
		} else {
			layers--
		}

		if S[i] == '(' && S[i+1] == ')' {
			res += 1 << uint(layers-1)
		}
	}
	return res
}

func main() {

	x := "(()(()))"
	fmt.Println(scoreOfParentheses(x))
}
