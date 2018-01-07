package main

import "math"

func main() {
}

//Very hard ;)
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return int(math.Max(float64(len(a)), float64(len(b))))
}
