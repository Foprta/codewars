package main

import (
	"reflect"
	"testing"
)

func TestRace(t *testing.T) {
	type args struct {
		v1 int
		v2 int
		g  int
	}
	tests := []struct {
		name string
		args args
		want [3]int
	}{
		{"720, 850, 70 = {0, 32, 18}", args{720, 850, 70}, [3]int{0, 32, 18}},
		{"820, 81, 550 = {-1, -1, -1}", args{820, 81, 550}, [3]int{-1, -1, -1}},
		{"80, 91, 37 = {3, 21, 49}", args{80, 91, 37}, [3]int{3, 21, 49}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Race(tt.args.v1, tt.args.v2, tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Race() = %v, want %v", got, tt.want)
			}
		})
	}
}
