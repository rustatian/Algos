package main

func main() {
	print(RepeatStr(4, "a"))
}

// Write a function called repeatStr which repeats the given string string exactly n times.
func RepeatStr(repititions int, value string) string {
	var newStr string
	for i := 0; i <= repititions-1; i++ {
		newStr += value
	}
	return newStr
}
