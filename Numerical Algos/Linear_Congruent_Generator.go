package main

import (
	"fmt"
	"flag"
)

//X(n+1) = (A x X(n) + B) % M, where A, B, M are constant
//X(0) - is initial number
//https://ru.wikipedia.org/wiki/Линейный_конгруэнтный_метод

func main() {
	var (
		A = flag.Int("A", 7, "A value")
		B = flag.Int("B", 5, "B value")
		M = flag.Int("M", 11, "M value")
		numIterations = flag.Int("iterNum", 10, "Iterations count")
	)

	var Xn1 int
	var Xn int = 0

	for i := 0; i <= *numIterations; i ++ {
		Xn1 = ((*A) * (Xn) + (*B)) % (*M)
		Xn = Xn1
	}

	fmt.Printf("You number after %d iterations is: %d", *numIterations, Xn1)
}

