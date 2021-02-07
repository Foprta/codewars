package main

import (
	"fmt"
	"strings"
)

// Accum https://www.codewars.com/kata/5544c7a5cb454edb3c000047/train/go
func Accum(s string) string {
	result := ""

	for idx, symbol := range strings.Split(s, "") {
		if idx > 0 {
			result += "-"
		}
		result += strings.ToUpper(symbol) + strings.Repeat(strings.ToLower(symbol), idx)
	}

	return result
}

func main() {
	fmt.Print(Accum("sd"))
}
