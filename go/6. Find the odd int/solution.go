package main

// FindOdd https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go
func FindOdd(seq []int) int {
	counter := make(map[int]int)

	for _, num := range seq {
		counter[num]++
	}

	for key, val := range counter {
		if val%2 == 1 {
			return key
		}
	}

	return 0
}

func main() {
	FindOdd([]int{1, 3, 5, 5, 3, 6})
}
