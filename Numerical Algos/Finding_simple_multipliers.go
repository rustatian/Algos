package main

import (
	"fmt"
	"math"
)

//We can solve that problem by 2 ways: just iterate or use some rules

func main() {

}

func smartSolution() {

	//1. All even numbers can be divided by 2, so we must divide by 2 and odd numbers. So, time reduced twice
	//2. We must check multipliers only to sqrt of num, n = p x q, so p and q must be equal or less sqrt(n)

	// Smart algo O(sqrt(n))
	num := 1322341234124
	i := 2
	var multipliers []int

	// First - check dividing by 2
	for num % 2 == 0 {
		multipliers = append(multipliers, num)
		num = num / 2
	}

	// Scan for odd multipliers
	i = 3
	max := math.Sqrt(float64(num))

	for float64(i) <= max {
		for num % i == 0 {
			multipliers = append(multipliers, num)
			num = num / i

			// New upper bound
			max = math.Sqrt(float64(num))
		}
		i = i + 2
	}

	if num > 1 {
		multipliers = append(multipliers, num)
	}

	fmt.Println(multipliers)
}



func simpleSolution(){
	num := 131
	i := 2
	var multipliers []int

	for i < num {
		// number Mod i == 0
		for (num % i) == 0 {
			multipliers = append(multipliers, num)
			num = num / i
		}
		i ++
	}

	if num > 1 {
		multipliers = append(multipliers, num)
	}

	fmt.Println(multipliers)
}
