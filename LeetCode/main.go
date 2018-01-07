package main

import (
	"fmt"
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
}
