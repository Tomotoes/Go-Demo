package main

import "sort"

func heightChecker(heights []int) int {

	sort.Ints(heights)
	count := 0
	for i, v := range heights {
		if former[i] != v {
			count++
		}
	}
	return count
}
