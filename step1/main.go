package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println(Solution([]int{3,4,3,0,2,2,3,0,0}))
	fmt.Println(Solution([]int{4,2,0}))
	fmt.Println(Solution([]int{4,4,3,3,1,0}))
}

func Solution(ranks []int) int {
	// write your code in Go 1.4
	sort.Ints(ranks)
	report := 0
	for i,_ := range ranks {
		if Superior(ranks, ranks[i] + 1) {
			report = report + 1
		}
	}
	return report
}

func Superior(list []int, x int) bool {
	for _, val := range list {
		if x == val {
			return true
			continue
		}
	}
	return false
}
