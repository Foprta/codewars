package main

import "testing"

func TestDivide(t *testing.T) {
	type args struct {
		weight int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"4", args{4}, true},
		{"2", args{2}, false},
		{"11", args{11}, false},
		{"5", args{5}, false},
		{"88", args{88}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.args.weight); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
