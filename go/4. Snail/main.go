package main

// Snail https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1/train/go
func Snail(snaipMap [][]int) []int {
	result := []int{}
	lines := len(snaipMap[0])*2 - 1
	moves := len(snaipMap[0])
	dx, dy := 1, 0
	x, y := -1, 0

	for ; lines > 0; lines-- {
		for range make([]int, moves) {
			x += dx
			y += dy
			result = append(result, snaipMap[y][x])
		}
		moves -= 1 & dx
		dx, dy = -dy, dx
	}

	return result
}

func main() {
}
