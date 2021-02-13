package main

// GetGrade https://www.codewars.com/kata/55cbd4ba903825f7970000f5/train/go
func GetGrade(a, b, c int) rune {
	mean := (a + b + c) / 3

	if mean >= 90 {
		return 'A'
	} else if mean >= 80 {
		return 'B'
	} else if mean >= 70 {
		return 'C'
	} else if mean >= 60 {
		return 'D'
	}
	return 'F'
}

func main() {
}
