package main

// TwoSum https://www.codewars.com/kata/52c31f8e6605bcc646000082/train/go
func TwoSum(numbers []int, target int) [2]int {
	numbersIndexes := make(map[int]int)

	for idx, number := range numbers {
		if pairIdx, ok := numbersIndexes[target-number]; ok == true {
			return [2]int{pairIdx, idx}
		}
		numbersIndexes[number] = idx
	}

	return [2]int{0, 0}
}

func main() {
	TwoSum([]int{1, 3, 2}, 4)
}
