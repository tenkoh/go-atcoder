package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `4 2
3 4 5 4
1 4
`
	s2 := `3 3
1 1 1
1000000000 1000000000 1000000000
`
	s3 := `7 3
2 6 8 9 5 1 11
3 5 7
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
		{"s1", args{r: strings.NewReader(s1)}, 7},
		{"s2", args{r: strings.NewReader(s2)}, -1},
		{"s3", args{r: strings.NewReader(s3)}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
