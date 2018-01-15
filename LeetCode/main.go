package main

import (
	"fmt"
	"strconv"
)

//Stack implementation (because can't import stack in leetcode thought import directive
type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

type Queue []int

func (q Queue) Empty() bool { return len(q) == 0 }
func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}
func (q *Queue) Dequeue() int {
	v := (*q)[0]
	*q = (*q)[1:len(*q)]
	return v
}

func main() {
	//fmt.Println(anagramMappings([]int{40, 40}, []int{40, 40}))
}

func anagramMappings(A []int, B []int) []int {
	var bb []int

	for _, elem := range A {
		for j := 0; j < len(B); j++ {
			if elem == B[j] {
				bb = append(bb, j)
				break
			}
		}
	}
	return bb
}
