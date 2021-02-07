package kata

import "testing"

func TestBouncingBall(t *testing.T) {
	type args struct {
		h      float64
		bounce float64
		window float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"testequal(3, 0.66, 1.5, 3)", args{h: 3, bounce: 0.66, window: 1.5}, 3},
		{"testequal(40, 0.4, 10, 3)", args{h: 40, bounce: 0.4, window: 10}, 3},
		{"testequal(10, 0.6, 10, -1)", args{h: 10, bounce: 0.6, window: 10}, -1},
		{"testequal(40, 1, 10, -1)", args{h: 40, bounce: 1, window: 10}, -1},
		{"testequal(5, -1, 1.5, -1)", args{h: 5, bounce: -1, window: 1.5}, -1},
		{"testequal(10, 0.9, 9, 3)", args{h: 10, bounce: 0.9, window: 9}, 1},
		{"testequal(10, 0.9, 9.5, 1)", args{h: 10, bounce: 0.9, window: 9.5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BouncingBall(tt.args.h, tt.args.bounce, tt.args.window); got != tt.want {
				t.Errorf("BouncingBall() = %v, want %v", got, tt.want)
			}
		})
	}
}
