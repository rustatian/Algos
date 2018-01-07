package main

import "strings"

func main() {
}

func reverseWords(s string) string {
	var newsrt string

	for i := 1; i <= len(s); i++ {
		newsrt += string(s[len(s)-i])
	}
	newsrt += " "

	newsrt = strings.TrimRight(newsrt, " ")

	return newsrt
}
