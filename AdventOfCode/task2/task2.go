package main

import (
	"bufio"
	"fmt"
	"os"
)

type countingStruct struct {
	Id     int
	number string
	freq   int
}

// --> []countingStruct {row:1, num: 2, freq: 4}

func main() {
	file, _ := os.Open("/Users/0xdev/Projects/repo/GoAlgos/AdventOfCode/task2/task2")
	b := bufio.NewScanner(file)

	m := make(map[int]int, 100)
	//resSl := []int{}

	for b.Scan() {
		text := b.Text()
		temp := make(map[string]int, len(text))

		for i := 0; i < len(text); i ++ {
			if v, ok := temp[string(text[i])]; ok {
				temp[string(text[i])] = v + 1
			} else {
				temp[string(text[i])] = 1
			}
		}

		mm := mapUnique(temp)

		for k, v := range mm {
			if v == true {
				if _, ok := m[k]; ok {
					m[k] += 1
				} else {
					m[k] = 1
				}

			}
		}
		temp = nil
	}

	fmt.Print(25*249)

	fmt.Print(m)

	//for _, v := range m {
	//	resSl = append(resSl, v)
	//}
	//
	//out := 1
	//for _, v := range res {
	//	out *= v
	//	//fmt.Println(k + "///" + string(v))
	//}
	//fmt.Println(out)

}

func mapUnique(mm map[string]int) map[int]bool {
	tempM := make(map[int]bool)
	for _, v := range mm {
		if v > 1 {
			tempM[v] = true
		}
	}
	return tempM
}
