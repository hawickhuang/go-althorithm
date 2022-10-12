package leetcode

import (
	"reflect"
	"testing"
)

func Test_repeatSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct{
		name string
		args args
		want bool
	}{
		{name: "return true", args: args{s: "abab"}, want: true},
		{name: "return false", args: args{s: "aba"}, want: false},
		{name: "return true", args: args{s: "abcabcabcabc"}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repeatSubstringPattern()=%v, want: %v", got, tt.want)
			}
		})
	}
}
