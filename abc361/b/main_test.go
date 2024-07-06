package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `0 0 0 4 5 6
2 3 4 5 6 7
`
	s2 := `0 0 0 2 2 2
0 0 2 2 2 4
`
	s3 := `0 0 0 1000 1000 1000
10 10 10 100 100 100
`
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"s1", args{strings.NewReader(s1)}, true},
		{"s2", args{strings.NewReader(s2)}, false},
		{"s3", args{strings.NewReader(s3)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
