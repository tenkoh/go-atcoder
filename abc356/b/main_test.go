package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `2 3
10 20 30
20 0 10
0 100 100
`
	s2 := `2 4
10 20 30 40
20 0 10 30
0 100 100 0
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
		{"s1", args{r: strings.NewReader(s1)}, true},
		{"s2", args{r: strings.NewReader(s2)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
