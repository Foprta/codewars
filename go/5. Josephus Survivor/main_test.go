package main

import "testing"

func TestJosephusSurvivor(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"6, 2 = 5", args{6, 2}, 5},
		{"11, 19 = 10", args{11, 19}, 10},
		{"7, 3 = 4", args{7, 3}, 4},
		{"40, 3 = 28", args{40, 3}, 28},
		{"14, 2 = 13", args{14, 2}, 13},
		{"100, 1 = 100", args{100, 1}, 100},
		{"300, 300 = 265", args{300, 300}, 265},
		{"562, 2379 = 146", args{562, 2379}, 146},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JosephusSurvivor(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("JosephusSurvivor() = %v, want %v", got, tt.want)
			}
		})
	}
}
