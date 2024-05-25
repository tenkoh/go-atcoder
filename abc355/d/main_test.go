package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `3
1 5
7 8
3 7
`
	s2 := `3
3 4
2 5
1 6
`
	s3 := `2
1 2
3 4
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
		{"s2", args{strings.NewReader(s2)}, 3},
		{"s3", args{strings.NewReader(s3)}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
