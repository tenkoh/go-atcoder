package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve2(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"ex1", args{strings.NewReader("3 2 + 4 1 - *")}, 15},
		{"ex2", args{strings.NewReader("3 2 1 + *")}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve2(tt.args.r); got != tt.want {
				t.Errorf("solve2() = %v, want %v", got, tt.want)
			}
		})
	}
}
