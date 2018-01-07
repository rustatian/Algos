package main

import (
	"fmt"
	"strconv"
)

func main() {
}

func selfDividingNumbers(left int, right int) []int {
	var nums []int
	for i := left; i <= right; i++ {
		strNum := strconv.Itoa(i)    //155
		for _, num := range strNum { //155 % 1, % 5, % 5
			digit, _ := strconv.Atoi(string(num))
			fmt.Println(digit)
			if digit == 0 {
				goto exit
			}
			if int(i)%digit != 0 {
				goto exit
			}

		}
		nums = append(nums, i)

	exit:
	}

	return nums
}
