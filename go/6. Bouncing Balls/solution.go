package kata

// BouncingBall https://www.codewars.com/kata/5544c7a5cb454edb3c000047/train/go
func BouncingBall(h, bounce, window float64) int {
	// Input check
	if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}

	// Logic
	result := 1
	for currentHeight := h * bounce; currentHeight > window; currentHeight = currentHeight * bounce {
		result += 2
	}
	return result
}
