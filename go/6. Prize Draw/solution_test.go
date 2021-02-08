package main

import "testing"

func TestNthRank(t *testing.T) {
	type args struct {
		st string
		we []int
		n  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"first test", args{"Elijah,Chloe,Elizabeth,Matthew,Natalie,Jayden", []int{1, 3, 5, 5, 3, 6}, 2}, "Matthew"},
		{"second test", args{"COLIN,AMANDBA,AMANDAB,CAROL,PauL,JOSEPH", []int{1, 4, 4, 5, 2, 1}, 4}, "PauL"},
		{"third test", args{"Addison,Jayden,Sofia,Michael,Andrew,Lily,Benjamin", []int{4, 2, 1, 4, 3, 1, 2}, 4}, "Benjamin"},
		{"empty string", args{"", []int{4, 2, 1, 4, 3, 1, 2}, 4}, "No participants"},
		{"n > participatns count", args{"Chloe", []int{4, 2, 1, 4, 3, 1, 2}, 2}, "Not enough participants"},
		{"alphabetical sort", args{"William,Willaim", []int{1, 1}, 1}, "Willaim"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NthRank(tt.args.st, tt.args.we, tt.args.n); got != tt.want {
				t.Errorf("NthRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
