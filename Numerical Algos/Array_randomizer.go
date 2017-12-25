package main

import "fmt"

// Array (and slice) randomizer with time O(N)
// We can use linear congruent generator to get value between 0 and max len of array

func main() {
	// Yes, we can generate in for loop, but... for test purposes :)
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	//In for loop
	//slice := getBigSlice()

	for i, elem := range slice {
		//We can't get number more than len of the slice in LCG; 124332 and 1437698 just numbers, you can choose any
		rand := lcg(i*124332, elem*1437698, len(slice))

		//Save value in i-th index
		temp := slice[i]
		//Assing with random value in slice
		slice[i] = slice[rand]
		//Change initial i-th value with rand
		slice[rand] = temp
	}

	fmt.Print(slice)
}


// Use linear congruent generator
func lcg(A, B, M int) int {
	var Xn = 44 //Seed
	var Xn1 int

	for i := 0; i < M; i++ {
		Xn1 = ((A)*(Xn) + (B)) % (M)
		Xn = Xn1
	}
	return Xn1
}

//Or big slice
func getBigSlice() []int {
	var bigSlice []int
	for i := 0; i <= 1000; i++ {
		bigSlice = append(bigSlice, i)
	}
	return bigSlice
}
