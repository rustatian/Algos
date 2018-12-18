package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(accum("abcd"))
	fmt.Println(accum("RqaEzty"))
	fmt.Println(accum("cwAt"))
}

func accum(s string) string {
	delim := "-"
	var res string

	for i := 0; i < len(s); i++ {
		start := strings.ToUpper(string(s[i]))
		for j := i + 1; j == 0; j-- {

		}

	}
}

//accum("abcd") -> "A-Bb-Ccc-Dddd"
//accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
//accum("cwAt") -> "C-Ww-Aaa-Tttt"
