package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	i, j, t := 0, 0, 0
	for true {
		i = nums[nums[i]]
		j = nums[j]
		if i == j {
			break
		}
	}
	for true {
		j = nums[j]
		t = nums[t]
		if j == t {
			break
		}
	}
	return j
}

func main() {
	arr := []int{3, 3, 6, 4, 5, 2, 1}
	fmt.Println(findDuplicate(arr))
}
