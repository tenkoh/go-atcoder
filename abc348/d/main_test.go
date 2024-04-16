package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `4 4
S...
#..#
#...
..#T
4
1 1 3
1 3 5
3 2 1
2 3 1
`
	s3 := `4 5
..#..
.S##.
.##T.
.....
3
3 1 5
1 2 3
2 2 1
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
		{"1", args{r: strings.NewReader(s1)}, true},
		{"3", args{r: strings.NewReader(s3)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
