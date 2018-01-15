package main

import "fmt"

func main() {
	fmt.Println(hasAlternatingBits(7))
}

func hasAlternatingBits(n int) bool {
	binary := fmt.Sprintf("%b", n)
	//binary := strconv.FormatInt(int64(n), 2)

	for i := 0; i < len(binary)-1; i++ {
		if string(binary[i]) == string(binary[i+1]) {
			return false
		}
	}

	return true
}
