package main

import "math"

// ListSquared https://www.codewars.com/kata/55aa075506463dac6600010d/train/go
func ListSquared(m, n int) [][]int {
	result := [][]int{}

	for ; m <= n; m++ {
		divisors := GetDivisors(m)
		divisorsSqrSum := DivisorsSquaredSum(divisors)
		sumSqrt := math.Sqrt(float64(divisorsSqrSum))
		if sumSqrt == float64(int(sumSqrt)) {
			result = append(result, []int{m, divisorsSqrSum})
		}
	}

	return result
}

// DivisorsSquaredSum return sum of squared divisers
func DivisorsSquaredSum(divs []int) int {
	sum := 0
	for _, n := range divs {
		sum += n * n
	}
	return sum
}

// GetDivisors return dividers of n
func GetDivisors(n int) []int {
	res := []int{}
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if n/i == i {
				res = append(res, i)
			} else {
				res = append(res, i, n/i)
			}
		}
	}
	return res
}

func main() {
}
