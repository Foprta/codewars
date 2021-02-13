package main

import (
	"reflect"
	"testing"
)

func TestPickPeaks(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want PosPeaks
	}{
		{"1, 3, 1", args{[]int{1, 3, 1}}, PosPeaks{Pos: []int{1}, Peaks: []int{3}}},
		{"1, 3, 3, 1", args{[]int{1, 3, 3, 1}}, PosPeaks{Pos: []int{1}, Peaks: []int{3}}},
		{"1, 3, 3, 1, 4, 2", args{[]int{1, 3, 3, 1, 4, 2}}, PosPeaks{Pos: []int{1, 4}, Peaks: []int{3, 4}}},
		{"1, 3, 3, 3, 4, 2", args{[]int{1, 3, 3, 3, 4, 2}}, PosPeaks{Pos: []int{4}, Peaks: []int{4}}},
		{"2, 2, 1, 2, 2", args{[]int{2, 2, 1, 2, 2}}, PosPeaks{Pos: []int{}, Peaks: []int{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PickPeaks(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PickPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}
