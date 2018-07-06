package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {

	return n > 0 && ((n & (n - 1)) == 0) //O(1)
}

func main() {

	for i := 0; i < 100000000000; i++ {
		if isPowerOfTwo(i) {
			fmt.Println(i)
		}

	}

}
