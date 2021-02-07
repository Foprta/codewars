package main

// NbDig https://www.codewars.com/kata/566fc12495810954b1000030/train/go
func NbDig(n int, d int) int {
	count := 0
	if d == 0 {
		count++
	}
	for i := 0; i <= n; i++ {
		count += numberOfDigits(d, int(i*i))
	}
	return count
}

// numberOfDigits counts number of given digits in given value
func numberOfDigits(digit, value int) int {
	count := 0
	for value != 0 {
		if value%10 == digit {
			count++
		}
		value /= 10
	}
	return count
}

func main() {
}
