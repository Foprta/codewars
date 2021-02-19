package main

import (
	"reflect"
	"testing"
)

func TestGetDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1 = 1", args{1}, []int{1}},
		{"2 = 1,2", args{2}, []int{1, 2}},
		{"4 = 1,2,4", args{4}, []int{1, 4, 2}},
		{"8 = 1,2,4,8", args{8}, []int{1, 8, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDivisors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivisorsSquaredSum(t *testing.T) {
	type args struct {
		divs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1,2,3 = 14", args{[]int{1, 2, 3}}, 14},
		{"1,2,10 = 105", args{[]int{1, 2, 10}}, 105},
		{"1,2,10,20 = 505", args{[]int{1, 2, 10, 20}}, 505},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DivisorsSquaredSum(tt.args.divs); got != tt.want {
				t.Errorf("DivisorsSquaredSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListSquared(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1, 250", args{1, 250}, [][]int{{1, 1}, {42, 2500}, {246, 84100}}},
		{"250, 500", args{250, 500}, [][]int{{287, 84100}}},
		{"300, 600", args{300, 600}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListSquared(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSquared() = %v, want %v", got, tt.want)
			}
		})
	}
}
