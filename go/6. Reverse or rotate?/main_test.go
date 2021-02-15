package main

import "testing"

func TestRevrot(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"123456987654, 6", args{"123456987654", 6}, "234561876549"},
		{"123456987653, 6", args{"123456987653", 6}, "234561356789"},
		{"66443875, 4", args{"66443875", 4}, "44668753"},
		{"66443875, 8", args{"66443875", 8}, "64438756"},
		{"664438769, 8", args{"664438769", 8}, "67834466"},
		{"123456779, 8", args{"123456779", 8}, "23456771"},
		{", 8", args{"", 8}, ""},
		{"123456779, 0", args{"123456779", 0}, ""},
		{"563000655734469485, 4", args{"563000655734469485", 4}, "0365065073456944"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Revrot(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Revrot() = %v, want %v", got, tt.want)
			}
		})
	}
}
