package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Decompose https://www.codewars.com/kata/54f8693ea58bce689100065f/train/go
func Decompose(s string) []string {
	result := []string{}

	// Wrong inputs
	if len(s) < 3 {
		if s == "0" {
			return result
		}
		return []string{s}
	}

	// "0.66"-like inputs
	if strings.Contains(s, ".") {
		s = Normalize(s)
	}

	split := strings.Split(s, "/")
	numer, _ := strconv.Atoi(split[0])
	denomer, _ := strconv.Atoi(split[1])

	if numer >= denomer {
		result = append(result, fmt.Sprintf("%d", numer/denomer))
		numer = numer % denomer
	}

	for numer > 0 && denomer > 0 {
		tmp := int(math.Ceil(float64(denomer) / float64(numer)))
		result = append(result, fmt.Sprintf("1/%d", tmp))
		numer, denomer = pmod(-denomer, numer), denomer*tmp
	}

	return result
}

// Normalize returns fraction from float
func Normalize(s string) string {
	split := strings.Split(s, ".")
	integer, fraction := split[0], split[1]

	numer, denomer := "", "1"+strings.Repeat("0", len(fraction))

	if integer == "0" {
		numer = fmt.Sprintf("%s", fraction)
	} else {
		numer = fmt.Sprintf("%s%s", integer, fraction)
	}

	return fmt.Sprintf("%s/%s", numer, denomer)
}

// mod function with positive result
func pmod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	return x + d
}
