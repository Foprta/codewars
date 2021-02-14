package main

// JosephusSurvivor https://www.codewars.com/kata/555624b601231dc7a400017a/train/go
func JosephusSurvivor(n, k int) int {
	deleted := make(map[int]bool)

	for i := 0; ; {
		moves := (k-1)%(n-len(deleted)) + 1
		for moves > 0 || deleted[i] {
			i++
			if i > n {
				i = i % n
			}
			if !deleted[i] {
				moves--
			}
		}
		deleted[i] = true
		if len(deleted) == n {
			return i
		}
	}
}

func main() {
}
