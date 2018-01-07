package main

import "sort"

func main() {
}

func distributeCandies(candies []int) int {
	sort.Ints(candies)
	var count int = 1
	for i := 1; i < len(candies) && count < len(candies)/2; i++ {
		if candies[i] > candies[i-1] {
			count++
		}
	}
	return count
}
