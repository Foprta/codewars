package main

import "testing"

func TestAccum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ab", args{s: "ab"}, "A-Bb"},
		{"fgR", args{s: "fgR"}, "F-Gg-Rrr"},
		{"RqaEzty", args{s: "RqaEzty"}, "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Accum(tt.args.s); got != tt.want {
				t.Errorf("Accum() = %v, want %v", got, tt.want)
			}
		})
	}
}
