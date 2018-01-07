package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
}

func arrayPairSum(nums []int) int {
	nums = []int{1, 4, 3, 2}
	sort.Ints(nums)

	var sum float64

	for i := 0; i < len(nums)-1; i += 2 {
		sum += math.Min(float64(nums[i]), float64(nums[i+1]))
	}

	fmt.Println(sum)
	return 1
}
