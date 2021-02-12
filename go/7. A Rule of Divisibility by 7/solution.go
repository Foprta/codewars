package main

// Seven https://www.codewars.com/kata/55e6f5e58f7817808e00002e/train/go
func Seven(n int64) []int {
	res := []int{0, 0}

	for i := 1; n > 99; i++ {
		n = n/10 - (n%10)*2
		res = []int{int(n), i}
	}

	return res
}

func main() {
}
