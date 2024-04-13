package main

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		whole s
	}
	tests := []struct {
		name string
		args args
		want []s
	}{
		// TODO: Add test cases.
		{"1", args{whole: s{l: 3, r: 19}}, []s{{3, 4}, {4, 8}, {8, 16}, {16, 18}, {18, 19}}},
		{"2", args{whole: s{l: 0, r: 1024}}, []s{{0, 1024}}},
		{"3", args{whole: s{l: 6, r: 15}}, []s{{6, 8}, {8, 12}, {12, 14}, {14, 15}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.whole); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
