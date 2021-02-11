package main

// Race https://www.codewars.com/kata/55e2adece53b4cdcb900006c/train/go
func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	seconds := g * 3600 / (v2 - v1)

	hours := seconds / 3600
	seconds %= 3600
	minutes := seconds / 60
	seconds %= 60

	return [3]int{hours, minutes, seconds}
}

func main() {
	Race(720, 850, 70)
}
