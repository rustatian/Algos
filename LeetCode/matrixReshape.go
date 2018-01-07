package main

func main() {
}

func matrixReshape(nums [][]int, r int, c int) [][]int {

	res := [][]int{[]int{}, []int{}}
	var rows, cols int = 0, 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			res[rows][cols] = nums[i][j]
			cols++
			if cols == c {
				rows++
				cols = 0
			}
		}
	} //[1, 2, 3, 4]

	return res
}
