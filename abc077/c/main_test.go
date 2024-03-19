package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `2
1 5
2 4
3 6
`
	s2 := `3
1 1 1
2 2 2
3 3 3
`
	s3 := `6
3 14 159 2 6 53
58 9 79 323 84 6
2643 383 2 79 50 288
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
		{"1", args{r: strings.NewReader(s1)}, 3},
		{"2", args{r: strings.NewReader(s2)}, 27},
		{"3", args{r: strings.NewReader(s3)}, 87},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
