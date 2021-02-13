package main

import "testing"

func TestGetGrade(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"95, 90, 93", args{95, 90, 93}, 'A'},
		{"70, 70, 100", args{70, 70, 100}, 'B'},
		{"70, 70, 70", args{70, 70, 70}, 'C'},
		{"65, 70, 59", args{65, 70, 59}, 'D'},
		{"44, 55, 52", args{44, 55, 52}, 'F'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGrade(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("GetGrade() = %v, want %v", got, tt.want)
			}
		})
	}
}
