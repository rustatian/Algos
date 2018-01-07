package main

import "strconv"

func main() {
}

func findComplement(num int) int {
	binStr := strconv.FormatInt(int64(num), 2)

	var resStr string
	for _, s := range binStr {
		if string(s) == "0" {
			resStr += "1"
		} else {
			resStr += "0"
		}

	}

	i, _ := strconv.ParseInt(string(resStr), 2, 32)

	return int(i)
}
