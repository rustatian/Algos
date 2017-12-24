package main

import (
	"flag"
	"math"
	"fmt"
)

// A ^ 2xM = (A^M) ^ 2
// A^(MxN) = A^M x A^N
// That code compared just to 7x7x7x7x7x7x7 <- computations like this

func main() {
	var (
		A = flag.Float64("Float64", 7, "Number")
		P = flag.Int("Power", 52, "Power")
	)
	flag.Parse()

	fmt.Printf("Real result by internal func : %e", math.Pow(*A, float64(*P)))
	fmt.Println()
	fmt.Println("------------")

	var i float64 = 2
	var res float64 = 1


	for i <= float64(*P) {
		res = math.Pow(*A, float64(i))
		i = i * 2
	}
	//
	if (float64(*P) - i) != 0 {
		res = math.Pow(*A, float64(*P) - i/2) * res
	}

	if *P == 1 {
		res = *A
	}

	fmt.Printf("Calculated result : %e", res)

}
