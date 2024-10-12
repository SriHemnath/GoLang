package main

import "fmt"

func main() {
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{1, 2}, []int{2, 1}))
}

func canCompleteCircuit(gas, cost []int) int {
	start := 0
	total := 0
	tank := 0

	for i := 0; i < len(gas); i++ {
		tank = tank + gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			total = total + tank
			tank = 0
		}
	}
	if total+tank < 0 {
		return -1
	} else {
		return start
	}
}
