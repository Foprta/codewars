package main

import (
	"fmt"
	"strings"
)

// Solution https://www.codewars.com/kata/51ba717bb08c1cd60f00002f/train/go
func Solution(list []int) string {
	result := []string{}
	accum := 0

	for i := 1; i <= len(list); i++ {
		if len(list)-i > 0 && list[i]-list[i-1] == 1 {
			accum++
		} else {
			if accum == 0 {
				result = append(result, fmt.Sprint(list[i-1]))
			} else if accum == 1 {
				result = append(result, fmt.Sprint(list[i-1]-accum))
				result = append(result, fmt.Sprint(list[i-1]))
			} else {
				result = append(result, fmt.Sprintf("%d-%d", list[i-1]-accum, list[i-1]))
			}
			accum = 0
		}
	}

	return strings.Join(result, ",")
}

func main() {
}
