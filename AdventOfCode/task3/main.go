package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type countingStruct struct {
	Id     int
	number string
	freq   int
}

// --> []countingStruct {row:1, num: 2, freq: 4}

func main() {
	file, _ := os.Open("/Users/0xdev/Projects/repo/GoAlgos/AdventOfCode/task3/task.txt")
	b := bufio.NewScanner(file)

	m := make([]string, 0)

	for i := 0; b.Scan(); i ++ {
		text := b.Text()
		m = append(m, text)
	}

	res := make(map[int]int)
	m2 := make([]string, len(m), cap(m))

	copy(m2, m)
	sort.Strings(m)

	for i := 0; i < len(m)-1; i ++ {
		r1 := m[i] // get row 0
		r2 := m[i+1]

		// iterate and check
		for ij := 0; ij < len(r1); ij ++ {
			if r1[ij] != r2[ij] {
				res[i] ++
			}
		}
	}

	for k, v := range res {
		if v == 1 {
			fmt.Print(fmt.Sprintf("The slices are: %s and %s", m[k], m[k + 1]))
		}
	}
}
