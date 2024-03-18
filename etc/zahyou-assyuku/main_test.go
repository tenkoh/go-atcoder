package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		as []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"ex1", args{[]int{12, 43, 7, 15, 9}}, []int{2, 4, 0, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.as); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
