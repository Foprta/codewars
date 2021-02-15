package main

// Revrot https://www.codewars.com/kata/56b5afb4ed1f6d5fb0000991/train/go
func Revrot(s string, n int) string {
	if n == 0 || n > len(s) {
		return ""
	}

	result := ""

	for i := 0; i+n <= len(s); i += n {
		sum := 0
		tmp := s[i : i+n]

		for _, digit := range tmp {
			integer := int(digit) - 48
			sum += integer * integer * integer
		}

		if sum%2 == 0 {
			result += reverseString(tmp)
		} else {
			result += rotateString(tmp)
		}
	}

	return result
}

func reverseString(s string) string {
	b := make([]byte, len(s))
	var j int = len(s) - 1
	for i := 0; i <= j; i++ {
		b[j-i] = s[i]
	}

	return string(b)
}

func rotateString(s string) string {
	return s[1:len(s)] + s[0:1]
}

func main() {
	Revrot("213123123123123123123", 3)
}
