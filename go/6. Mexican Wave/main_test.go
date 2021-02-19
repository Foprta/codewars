package main

import (
	"reflect"
	"testing"
)

func Test_wave(t *testing.T) {
	type args struct {
		words string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"abc", args{"abc"}, []string{"Abc", "aBc", "abC"}},
		{" x yz", args{" x yz"}, []string{" X yz", " x Yz", " x yZ"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wave(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wave() = %v, want %v", got, tt.want)
			}
		})
	}
}
