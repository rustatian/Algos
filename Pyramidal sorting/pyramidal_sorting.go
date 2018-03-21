package main

import "fmt"

func siftDown(numbers *[]int, root, bottom int) {
	var maxChild int
	var done bool

	for (root*2 <= bottom) && (!done) {
		if root*2 == bottom {
			maxChild = root * 2
		} else if (*numbers)[root*2] > (*numbers)[root*2+1] {
			maxChild = root * 2
		} else {
			maxChild = root*2 + 1
		}

		if (*numbers)[root] < (*numbers)[maxChild] {
			temp := (*numbers)[root]
			(*numbers)[root] = (*numbers)[maxChild]
			(*numbers)[maxChild] = temp
			root = maxChild
		} else {
			done = true
		}
	}
}

func heapSort(numbers *[]int, arraySize int) {
	for i := arraySize/2 - 1; i >= 0; i-- {
		siftDown(&(*numbers), i, arraySize)
	}
	for i := arraySize - 1; i >= 1; i-- {
		temp := (*numbers)[0]
		(*numbers)[0] = (*numbers)[i]
		(*numbers)[i] = temp
		siftDown(&(*numbers), 0, i-1)
	}

}

func main() {
	a := []int{2, 23, 33, 44, 1, 2, 2, 2, 4, 6, 99}

	heapSort(&a, len(a)-1)

	for _, m := range a {
		fmt.Println(m)
	}
}
