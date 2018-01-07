package main

func main() {
}

func judgeCircle(moves string) bool {
	coord := [2]int{0, 0}

	for _, char := range moves {
		switch string(char) {
		case "R":
			coord[0] = coord[0] + 1
		case "L":
			coord[0] = coord[0] - 1
		case "U":
			coord[1] = coord[1] + 1
		case "D":
			coord[1] = coord[1] - 1
		}
	}

	if coord[0] == 0 && coord[1] == 0 {
		return true
	}
	return false
}
