package main

import "strings"

func main() {
}

func findWords(words []string) []string {
	var newW []string

	for i := 0; i < len(words); i++ {
		if len(words[i]) == 1 {
			newW = append(newW, words[i])
			continue
		}

		if strings.ContainsAny(words[i], "QWERTYUIOPqwertyuiop") && strings.ContainsAny(words[i], "ASDFGHJKLasdfghjkl") {
			continue
		}
		if strings.ContainsAny(words[i], "QWERTYUIOPqwertyuiop") && strings.ContainsAny(words[i], "ZXCVBNMzxcvbnm") {
			continue
		}
		if strings.ContainsAny(words[i], "ZXCVBNMzxcvbnm") && strings.ContainsAny(words[i], "ASDFGHJKLasdfghjkl") {
			continue
		}

		newW = append(newW, words[i])
	}

	return newW
}
