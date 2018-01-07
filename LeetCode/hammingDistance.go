package main

import (
	"strconv"
	"strings"
)

func main() {
}

func hammingDistance(x int, y int) int {
	return strings.Count(strconv.FormatInt(int64(x^y), 2), "1")
}
