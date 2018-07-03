package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {

	for i := range gas {
		var curr = i
		var gasAmount = gas[curr]
		for {
			next := curr + 1
			if next == len(gas) {
				next = 0
			}
			gasAmount -= cost[curr]

			if gasAmount < 0 {
				break
			}
			gasAmount += gas[next]
			curr++
			if curr == len(gas) {
				curr = 0
			}

			if curr == i {

				return i
			}

		}
	}

	return -1
}

func main() {

	gas1 := []int{1, 2, 3, 4, 5}
	cost1 := []int{3, 4, 5, 1, 2}

	gas2 := []int{2, 3, 4}
	cost2 := []int{3, 4, 3}

	gas3 := []int{5, 1, 2, 3, 4}
	cost3 := []int{4, 4, 1, 5, 1}

	fmt.Println(canCompleteCircuit(gas1, cost1))
	fmt.Println(canCompleteCircuit(gas2, cost2))
	fmt.Println(canCompleteCircuit(gas3, cost3))

}
