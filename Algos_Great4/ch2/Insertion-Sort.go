package main

import "fmt"

func main() {
	sl := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(insertionSortReverse(sl))
	fmt.Println(insertionSort(sl))
}

func insertionSort(A []int) []int {
	// at the beginning we have elements sorted A[0..j-1]
	// and A[j + 0..n] are not
	// elements A[0..j-1] are being at positions from 0 to j-1, but in different order. Now they are sorted. This property called - LOOP INVARIANT
	for j := 1; j < len(A); j++ {
		// getting j value
		key := A[j]
		// getting previous value and call it i
		i := j - 1
		// while (in golang it's for)
		// while i >=0 and i'th element bigger than j'th (key)
		// do the swap
		// move back (i = i - 1)
		for i >= 0 && A[i] >= key {
			A[i], A[i+1] = A[i+1], A[i]
			i--
		}
		A[i+1] = key // end property of INVARIANT
	}
	return A
}

// Task 2.1.2
// The same but in reverse order
func insertionSortReverse(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] <= key {
			A[i], A[i+1] = A[i+1], A[i]
			i--
		}
		A[i+1] = key
	}
	return A
}

// Task 2.1.3
// Pseudocode
// value n

/*
FOR i = 0, i < A.Length, i ++
	IF A[i] = n
		return i
	return <nil>
*/

// Task 2.1.4
/*




 */
