package main

import "sort"

func sortedSquares(A []int) []int {
	for i, v := range A {
		A[i] = v * v
	}
	sort.Ints(A)
	return A
}
