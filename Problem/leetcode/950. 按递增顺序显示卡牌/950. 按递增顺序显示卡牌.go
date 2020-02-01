package main

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	length := len(deck)
	half := length / 2
	r := make([]int, length)
	for i := 0; i < half; i++ {
	}
	return r
}
