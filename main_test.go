package main

import (
	"reflect"
	"testing"
)

var _ = func() bool {
	test = true
	return false
}()

func Test_merge(t *testing.T) {
	type args struct {
		arr [4]int
	}
	tests := []struct {
		name string
		args args
		want [4]int
	}{
		{
			name: "",
			args: args{arr: [4]int{4, 4, 8, 0}},
			want: [4]int{8, 8, 0, 0},
		},
		{
			name: "",
			args: args{arr: [4]int{2, 4, 4, 0}},
			want: [4]int{2, 8, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
