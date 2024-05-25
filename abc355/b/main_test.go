package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `3 2
3 2 5
4 1
`
	s2 := `3 2
3 1 5
4 2
`
	s3 := `1 1
1
2
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
		{"1", args{strings.NewReader(s1)}, true},
		{"2", args{strings.NewReader(s2)}, false},
		{"3", args{strings.NewReader(s3)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
