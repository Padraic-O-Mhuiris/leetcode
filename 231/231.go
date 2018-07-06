package main

import (
	"fmt"
	"strconv"
)

func isPowerOfTwo(n int) bool {

	if n <= 0 {
		return false
	}
	numstring := strconv.FormatInt(int64(n), 2)
	powflag := false

	for _, a := range numstring {
		if powflag && a == '1' {
			return false
		}

		if !powflag && a == '1' {
			powflag = true
		}
	}

	if !powflag {
		return false
	}

	return true
}

func main() {

	for i := 0; i < 100; i++ {
		fmt.Println(i, isPowerOfTwo(i))
	}

}
