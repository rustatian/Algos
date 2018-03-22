package main

// Given a year, return the century it is in.
//
// The first century spans from the year 1 up to and including the year 100,
// the second - from the year 101 up to and including the year 200, etc.

func main() {
	println(century(int(101)))
}

func century(year int) int {
	return (year + 99) / 100
}
