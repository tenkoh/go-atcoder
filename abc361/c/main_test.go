package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `5 2
3 1 5 4 9
`
	s2 := `6 5
1 1 1 1 1 1
`
	s3 := `8 3
31 43 26 6 18 36 22 13
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
		{"s1", args{strings.NewReader(s1)}, 2},
		{"s2", args{strings.NewReader(s2)}, 0},
		{"s3", args{strings.NewReader(s3)}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
