package main

import "testing"

func TestFindOdd(t *testing.T) {
	type args struct {
		seq []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first", args{[]int{1, 2, 2, 2, 1}}, 2},
		{"second", args{[]int{1, 2, 1, 2, 1}}, 1},
		{"third", args{[]int{1, 2, 2, 2, 1, 3, 3, 3, 3, 3, 3}}, 2},
		{"fourth", args{[]int{2, 2, 1, 6, 6, 8, 8, 3, 3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOdd(tt.args.seq); got != tt.want {
				t.Errorf("FindOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
