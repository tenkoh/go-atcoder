package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `4
1 -2 -1
`
	s2 := `3
0 0
`
	s3 := `6
10 20 30 40 50
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
		{"1", args{r: strings.NewReader(s1)}, 2},
		{"2", args{r: strings.NewReader(s2)}, 0},
		{"3", args{r: strings.NewReader(s3)}, -150},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
