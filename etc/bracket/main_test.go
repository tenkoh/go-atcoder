package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []string
	}{
		// TODO: Add test cases.
		{"ex1", args{strings.NewReader("(()(())())(()())")}, true, []string{"1 2", "4 5", "3 6", "7 8", "0 9", "11 12", "13 14", "10 15"}},
		{"ex1", args{strings.NewReader("(()(())())(()()")}, false, []string{"1 2", "4 5", "3 6", "7 8", "0 9", "11 12", "13 14"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := solve(tt.args.r)
			if got != tt.want {
				t.Errorf("solve() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("solve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
