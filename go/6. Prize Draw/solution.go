package main

import (
	"sort"
	"strings"
)

type rank struct {
	name  string
	score int
}

// NthRank https://www.codewars.com/kata/5616868c81a0f281e500005c/train/go
func NthRank(st string, we []int, n int) string {
	if len(st) == 0 {
		return "No participants"
	}

	names := strings.Split(st, ",")
	if n > len(names) {
		return "Not enough participants"
	}

	rankings := make([]rank, 0)
	for idx, name := range names {
		score := 0
		for _, symbol := range strings.ToLower(name) {
			score += int(rune(symbol)) - 96
		}
		rankings = append(rankings, rank{name, (score + len(name)) * we[idx]})
	}

	sort.Slice(rankings, func(i, j int) bool {
		if rankings[i].score == rankings[j].score {
			return rankings[i].name < rankings[j].name
		}
		return rankings[i].score > rankings[j].score
	})

	return rankings[n-1].name
}

func main() {
	NthRank("Elijah,Chloe,Elizabeth,Matthew,Natalie,Jayden", []int{1, 3, 5, 5, 3, 6}, 1)
}
