package main

import (
	"flag"
	"fmt"
)

// X(n+1) = (A x X(n) + B) % M, where A, B, M are constant
// X(0) - is initial number
// https://ru.wikipedia.org/wiki/Линейный_конгруэнтный_метод
func main() {
	var (
		A = flag.Int("A", 2121, "A value")
		B = flag.Int("B", 43424, "B value")
		M = flag.Int("M", 44444444, "M value")
	)

	var Xn int = 0 // if we hold A and B, Xn will be seed
	var Xn1 int

	for i := 0; i < *M; i++ {
		Xn1 = ((*A)*(Xn) + (*B)) % (*M)
		Xn = Xn1
	}

	fmt.Printf("You number after %d iterations is: %d", *M, Xn1)
}
