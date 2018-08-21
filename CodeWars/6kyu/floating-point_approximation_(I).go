package main

import (
	"fmt"
	"math"
)

func main() {
	res := (math.Sqrt(1+1e-15*2) - 1) / 2
	fmt.Println(res)
}
