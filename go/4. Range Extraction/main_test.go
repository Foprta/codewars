package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]int{1}}, "1"},
		{"1,2", args{[]int{1, 2}}, "1,2"},
		{"1,2,3", args{[]int{1, 2, 3}}, "1-3"},
		{"-3,1,2,3,5,6,7", args{[]int{-3, 1, 2, 3, 5, 6, 7}}, "-3,1-3,5-7"},
		{"-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20", args{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}}, "-6,-3-1,3-5,7-11,14,15,17-20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.list); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
