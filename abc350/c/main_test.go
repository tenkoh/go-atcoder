package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `5
3 4 1 2 5`
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []*pair
	}{
		// TODO: Add test cases.
		{"s1", args{strings.NewReader(s1)}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
