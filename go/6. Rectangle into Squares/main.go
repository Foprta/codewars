package main

// SquaresInRect https://www.codewars.com/kata/55466989aeecab5aac00003e/train/go
func SquaresInRect(lng int, wdth int) []int {
	if lng == wdth {
		return nil
	}

	result := []int{}

	for lng != wdth {
		if lng > wdth {
			result = append(result, wdth)
			lng -= wdth
		} else {
			result = append(result, lng)
			wdth -= lng
		}
	}

	result = append(result, wdth)

	return result
}

func main() {
}
