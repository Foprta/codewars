package main

import (
	"reflect"
	"testing"
)

func TestSeven(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"477557101 = 28, 7", args{477557101}, []int{28, 7}},
		{"477557102 = 47, 7", args{477557102}, []int{47, 7}},
		{"1603 = 7, 2", args{1603}, []int{7, 2}},
		{"371 = 35, 1", args{371}, []int{35, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Seven(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seven() = %v, want %v", got, tt.want)
			}
		})
	}
}
