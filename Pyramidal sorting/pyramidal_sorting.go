package main

import "fmt"

func siftDown(numbers *[]int, root, bottom int) {
	n := *numbers

	var maxChild int
	var done bool

	for (root*2 <= bottom) && (!done) {
		if root*2 == bottom {
			maxChild = root * 2
		} else if n[root*2] > n[root*2+1] {
			maxChild = root * 2
		} else {
			maxChild = root*2 + 1
		}

		if n[root] < n[maxChild] {
			temp := n[root]
			n[root] = n[maxChild]
			n[maxChild] = temp
			root = maxChild
		} else {
			done = true
		}
	}
}

func heapSort(numbers *[]int, arraySize int) {
	n := *numbers
	for i := arraySize/2 - 1; i >= 0; i-- {
		siftDown(&n, i, arraySize)
	}
	for i := arraySize - 1; i >= 1; i-- {
		temp := n[0]
		n[0] = n[i]
		n[i] = temp
		siftDown(&n, 0, i-1)
	}

}

func main() {
	a := []int{2, 23, 33, 44, 1, 2, 2, 2, 4, 6, 99}

	heapSort(&a, len(a)-1)

	for _, m := range a {
		fmt.Println(m)
	}

}
