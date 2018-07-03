// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	numsMap := make(map[int]int) // map[key-type]value-type

	for i2, v := range nums {

		newTarget := target - v
		// Using indices as values and array values as keys

		if i1, ok := numsMap[newTarget]; ok {
			return []int{i1, i2}
		}
		numsMap[v] = i2
	}

	return []int{}
}

func main() {

	var arr = []int{2, 4, 12, 7, 2, 3, 4, 9}

	ans := twoSum(arr, 21)
	fmt.Println(ans)
}
