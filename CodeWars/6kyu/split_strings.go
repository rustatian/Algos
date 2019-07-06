package main

//Complete the solution so that it splits the string into pairs of two characters. If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').
//
//Examples:
//
//Solution("abc") //should return ["ab", "c_"]
//Solution("abcdef") //should return ["ab", "cd", "ef"]
func main() {
	println(Solution("awsaws"))
}

func Solution(str string) []string {
	var res []string
	if len(str)%2 == 0 {
		for i := 0; i < len(str)-1; i++ {
			res = append(res, string(str[i])+string(str[i+1]))
			i++
		}
	} else {
		sttr2 := str + "_"

		for i := 0; i < len(sttr2)-1; i++ {
			res = append(res, string(sttr2[i])+string(sttr2[i+1]))
			i++

		}
	}

	return res
}
