package main

// MaximumSubarraySum https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c/train/go
func MaximumSubarraySum(numbers []int) int {
	max := 0
	current := 0
	for _, num := range numbers {
		current += num
		if num > current {
			current = num
		}
		if current > max {
			max = current
		}
	}
	return max
}

func main() {
}
