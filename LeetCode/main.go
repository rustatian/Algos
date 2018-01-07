package main

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

type Queue []TreeNode

func (q Queue) Empty() bool { return len(q) == 0 }
func (q *Queue) Enqueue(v TreeNode) {
	*q = append(*q, v)
}
func (q *Queue) Dequeue() TreeNode {
	v := (*q)[0]
	*q = (*q)[1:len(*q)]
	return v
}

func main() {

}
