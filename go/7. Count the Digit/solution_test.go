package main

import (
	"testing"
)

func TestNbDig(t *testing.T) {
	type args struct {
		n int
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 in 10", args{n: 10, d: 1}, 4},
		{"5 in 550", args{n: 550, d: 5}, 213},
		{"0 in 4", args{n: 4, d: 0}, 1},
		{"0 in 5750", args{n: 5750, d: 0}, 4700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NbDig(tt.args.n, tt.args.d); got != tt.want {
				t.Errorf("NbDig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberOfDigits(t *testing.T) {
	type args struct {
		digit int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 in 1234", args{digit: 1, value: 1234}, 1},
		{"1 in 12223114", args{digit: 1, value: 12223114}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfDigits(tt.args.digit, tt.args.value); got != tt.want {
				t.Errorf("NumberOfDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
