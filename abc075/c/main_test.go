package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `7 7
1 3
2 7
3 4
4 5
4 6
5 6
6 7
`
	s2 := `3 3
1 2
1 3
2 3
`
	s3 := `6 5
1 2
2 3
3 4
4 5
5 6
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
		{"ex1", args{r: strings.NewReader(s1)}, 4},
		{"ex2", args{r: strings.NewReader(s2)}, 0},
		{"ex3", args{r: strings.NewReader(s3)}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
