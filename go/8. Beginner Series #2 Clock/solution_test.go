package kata

import "testing"

func TestPast(t *testing.T) {
	type args struct {
		h int
		m int
		s int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"01:00:00", args{h: 1, m: 0, s: 0}, 3600000},
		{"02:00:00", args{h: 2, m: 0, s: 0}, 7200000},
		{"00:10:10", args{h: 0, m: 10, s: 10}, 610000},
		{"00:01:01", args{h: 0, m: 1, s: 1}, 61000},
		{"01:01:01", args{h: 1, m: 1, s: 1}, 3661000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Past(tt.args.h, tt.args.m, tt.args.s); got != tt.want {
				t.Errorf("Past() = %v, want %v", got, tt.want)
			}
		})
	}
}
