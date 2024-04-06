package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `4
100 1
20 5
30 5
40 1
`
	s2 := `10
68 3
17 2
99 2
92 4
82 4
10 3
100 2
78 1
3 1
35 4
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
		{"1", args{strings.NewReader(s1)}, 40},
		{"2", args{strings.NewReader(s2)}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
