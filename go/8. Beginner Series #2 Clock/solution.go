package kata

// Past https://www.codewars.com/kata/55f9bca8ecaa9eac7100004a/train/go
func Past(h, m, s int) int {
	return (((h*60)+m)*60 + s) * 1000
}
