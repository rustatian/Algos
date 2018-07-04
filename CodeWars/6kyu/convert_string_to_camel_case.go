package main

import (
	"fmt"
)

// Complete the method/function so that it converts dash/underscore delimited words into camel casing.
// The first word within the output should be capitalized only if the original word was capitalized.

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior")) // returns "theStealthWarrior"
	fmt.Println(ToCamelCase("The_Stealth_Warrior")) // returns "TheStealthWarrior"
}

func ToCamelCase(s string) string {

	if s == "" {
		return ""
	}
	var b []byte
	b = append(b, s[0])
	for i := 1; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
			b = append(b, c)
		} else if c >= 'a' && c <= 'z' {
			b = append(b, c)
			continue
		} else {
			c := s[i+1]
			if c >= 'a' && c <= 'z' {
				c -= 'a' - 'A'
			}
			b = append(b, c)
			i += 1
		}
	}
	return string(b)
}
