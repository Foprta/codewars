package kata

import "testing"

func TestEvenOrOdd(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"10 is even", args{number: 10}, "Even"},
		{"11 is odd", args{number: 11}, "Odd"},
		{"-11 is odd", args{number: -11}, "Odd"},
		{"-10 is even", args{number: -10}, "Even"},
		{"0 is even", args{number: 0}, "Even"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvenOrOdd(tt.args.number); got != tt.want {
				t.Errorf("EvenOrOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
