package main

// Divide https://www.codewars.com/kata/55192f4ecd82ff826900089e/train/go
func Divide(weight int) bool {
	return weight > 2 && (weight-2)%2 == 0
}

func main() {
	Divide(7)
}
