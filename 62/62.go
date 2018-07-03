// Given an MxN grid where the user is positioned in the top right
// How many distinct paths are there that navigate to the bottom-right
// while only moving down and to the right

/*
Solution
------------
The problem can be easily solved by calculating the permutations for
each grids paths. A 7x3 grid will require 8 movements, 6 to the right,
2 moves down.
Therfore we can visualise the permutations of movements as:

	D D R R R R R R
	D R D R R R R R
	...
	...
	and so on

*/
package main

import (
	"fmt"
	"math/big"
)

func factorial(n *big.Int) (result *big.Int) {
	result = new(big.Int)

	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		result.SetInt64(1)
	default:
		result.Set(n)
		var one big.Int
		one.SetInt64(1)
		result.Mul(result, factorial(n.Sub(n, &one)))
	}
	return
}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	// let m always be the larger
	if m <= n {
		m, n = n, m
	}

	m, n = m-1, n-1

	a := factorial(big.NewInt(int64(m)))
	b := factorial(big.NewInt(int64(n)))
	c := factorial(big.NewInt(int64(m + n)))

	res := c.Div(c, a.Mul(a, b))

	return int(res.Int64())
}

func main() {

	ans := uniquePaths(2, 100)
	fmt.Println(ans)
}
