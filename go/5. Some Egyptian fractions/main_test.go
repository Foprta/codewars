package main

import (
	"reflect"
	"testing"
)

func TestDecompose(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"21/23", args{"21/23"}, []string{"1/2", "1/3", "1/13", "1/359", "1/644046"}},
		{"12/4", args{"12/4"}, []string{"3"}},
		{"5/2", args{"5/2"}, []string{"2", "1/2"}},
		{"3/2", args{"3/2"}, []string{"1", "1/2"}},
		{"1/2", args{"1/2"}, []string{"1/2"}},
		{"0.66", args{"0.66"}, []string{"1/2", "1/7", "1/59", "1/5163", "1/53307975"}},
		{"1.25", args{"1.25"}, []string{"1", "1/4"}},
		{"0", args{"0"}, []string{}},
		{"1", args{"1"}, []string{"1"}},
		{"wiki 7/15", args{"7/15"}, []string{"1/3", "1/8", "1/120"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decompose(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decompose(%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0.66", args{"0.66"}, "66/100"},
		{"0.2", args{"0.2"}, "2/10"},
		{"0.156", args{"0.156"}, "156/1000"},
		{"1.156", args{"1.156"}, "1156/1000"},
		{"3.1", args{"3.1"}, "31/10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Normalize(tt.args.s); got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pmod(t *testing.T) {
	type args struct {
		x int
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"-14 mod 6", args{-14, 6}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pmod(tt.args.x, tt.args.d); got != tt.want {
				t.Errorf("pmod() = %v, want %v", got, tt.want)
			}
		})
	}
}
