package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `3 5
5 1 8 9 7
`
	s2 := `3 5
4 2 9 7 5
`
	s3 := `4 12
13 9 6 5 2 7 16 14 8 3 10 11
`
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"s1", args{strings.NewReader(s1)}, 4},
		{"s2", args{strings.NewReader(s2)}, -1},
		{"s3", args{strings.NewReader(s3)}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
