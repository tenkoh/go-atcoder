package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `4
0 0
2 4
5 0
3 4
`
	s2 := `6
3 2
1 6
4 5
1 3
5 5
9 8
`
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{strings.NewReader(s1)}, []int{3, 3, 1, 1}},
		{"2", args{strings.NewReader(s2)}, []int{6, 6, 6, 6, 6, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
