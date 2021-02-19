package main

import "strings"

// wave https://www.codewars.com/kata/58f5c63f1e26ecda7e000029/train/go
func wave(words string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		if string(words[i]) == string(" ") {
			continue
		}

		result = append(result, words[:i]+strings.ToUpper(words[i:i+1])+words[i+1:])
	}

	return result
}

func main() {
}
