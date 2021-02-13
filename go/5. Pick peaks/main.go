package main

// PosPeaks type
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// PickPeaks https://www.codewars.com/kata/5279f6fe5ab7f447890006a7/train/go
func PickPeaks(array []int) PosPeaks {
	result := PosPeaks{Pos: make([]int, 0), Peaks: make([]int, 0)}

	for i := 1; i < len(array)-1; i++ {
		if array[i-1] >= array[i] || array[i+1] > array[i] {
			continue
		}

		if array[i] == array[i+1] {
			for j := i + 1; j < len(array)-1 && array[i] == array[j]; j++ {
				if array[j] > array[j+1] {
					result.Peaks = append(result.Peaks, array[i])
					result.Pos = append(result.Pos, i)
				}
			}
		} else {
			result.Peaks = append(result.Peaks, array[i])
			result.Pos = append(result.Pos, i)
		}

	}

	return result
}

func main() {
}
