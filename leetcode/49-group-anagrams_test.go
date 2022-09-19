package leetcode

import (
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{strs: []string{"eat","tea","tan","ate","nat","bat"}}, want: [][]string{[]string{"bat"},[]string{"nat","tan"},[]string{"ate","eat","tea"}}},
		{name: "test2", args: args{strs: []string{""}}, want: [][]string{[]string{""}}},
		{name: "test3", args: args{strs: []string{"a"}}, want: [][]string{[]string{"a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !comp2DSlice(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
