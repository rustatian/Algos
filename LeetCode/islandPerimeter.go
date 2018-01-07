package main

func main() {
}

//https://leetcode.com/problems/island-perimeter/
func islandPerimeter(grid [][]int) int {
	var island, neigh int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				island++
				if i < len(grid)-1 && grid[i+1][j] == 1 {
					neigh++
				}
				if j < len(grid[i])-1 && grid[i][j+1] == 1 {
					neigh++
				}
			}

		}
	}
	return island*4 - neigh*2

}
