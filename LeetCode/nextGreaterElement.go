package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{3, 1, 5, 7, 9, 2, 6}, []int{1, 2, 3, 5, 6, 7, 9, 11})) // res - [5,2,6,9,11,3,7]
}

//https://leetcode.com/problems/next-greater-element-i/
func nextGreaterElement(findNums []int, nums []int) []int {
	var hm map[int]int = make(map[int]int)
	var st Stack

	for _, nm := range nums {
		for !st.IsEmpty() && st.Peek() < nm {
			hm[st.Pop()] = nm
		}
		st.Push(nm)
	}
	for i := 0; i < len(findNums); i++ {
		_, ok := hm[findNums[i]]
		if ok {
			findNums[i] = hm[findNums[i]]
		} else {
			findNums[i] = -1
		}
	}

	return findNums
}
