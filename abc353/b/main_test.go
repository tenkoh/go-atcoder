package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `7 6
2 5 1 4 1 2 3
`
	s2 := `7 10
1 10 1 10 1 10 1
`
	s3 := `15 100
73 8 55 26 97 48 37 47 35 55 5 17 62 2 60
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
		{"s1", args{r: strings.NewReader(s1)}, 4},
		{"s2", args{r: strings.NewReader(s2)}, 7},
		{"s3", args{r: strings.NewReader(s3)}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
