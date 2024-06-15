package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `4
3 2 5 2
`
	s2 := `3
4 3 2
`
	s3 := `7
10 5 10 2 10 13 15
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
		{"s1", args{r: strings.NewReader(s1)}, 3},
		{"s2", args{r: strings.NewReader(s2)}, -1},
		{"s3", args{r: strings.NewReader(s3)}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
