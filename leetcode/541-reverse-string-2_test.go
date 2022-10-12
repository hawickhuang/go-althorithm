package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseString2(t *testing.T)  {
	type args struct {
		s string
		k int
	}

	tests := []struct{
		name string
		args args
		want string
	}{
		{name: "return bacdfeg", args: args{s: "abcdefg", k: 2}, want: "bacdfeg"},
		{name: "return bacd", args: args{s: "abcd", k: 2}, want: "bacd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString2(tt.args.s, tt.args.k)	; !reflect.DeepEqual(tt.want, got) {
				t.Errorf("reverseString2()=%v, want: %v", got, tt.want)
			}
		})
	}
}
