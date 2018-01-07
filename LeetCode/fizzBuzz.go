package main

import "strconv"

func main() {
}

func fizzBuzz(n int) []string {
	var out []string

	for i := 1; i <= n; i++ {
		switch i%3 == 0 {
		case true:
			switch i%5 == 0 {
			case true:
				out = append(out, "FizzBuzz")
			case false:
				out = append(out, "Fizz")
			}

		case false:
			switch i%5 == 0 {
			case true:
				out = append(out, "Buzz")
			case false:
				out = append(out, strconv.Itoa(i))
			}
		}
	}

	return out

}
