package main

import (
	"fmt"
	"math"
)

func main() {
	
}

// k-prime number
func CountKprimes(k, start, nd int)  []int {
	// n * 2^n - 1

	maxNumber := nd

	// maxNumber + 1
	isComposite := [1000001]bool{}

	// except all numbers which multiply 2
	for i := 4; i <= int(maxNumber); {
		isComposite[i] = true
		i += 2
	}

	// except all numbers which multiply founded prime numbers
	var nextPrime int = 3
	stopAt := math.Sqrt(float64(maxNumber))

	for nextPrime <= int(stopAt) {
		// except all numbers which multiply current prime number
		for i := nextPrime * 2; i <= maxNumber; i += nextPrime {
			isComposite[int(i)] = true
		}

		nextPrime = nextPrime + 2

		// goes to the next prime number except founded odd nums
		for nextPrime <= maxNumber && isComposite[int(nextPrime)] {
			nextPrime = nextPrime + 2
		}
	}

	var resSlice []int

	// Just sieve and get only prime numbers
	for i := 2; i <= int(maxNumber); i++ {
		if !isComposite[i] {
			resSlice = append(resSlice, i)
		}
	}


	return resSlice




	// your code
}


func Erythrophene_sieve() {



}

func Puzzle(s int) int {
	// your code
}
