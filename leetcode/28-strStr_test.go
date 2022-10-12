package leetcode

import (
	"reflect"
	"testing"
)

func Test_strStr(t *testing.T)  {
	type args struct {
		haystack string
		needle string
	}

	tests := []struct{
		name string
		args args
		want int
	}{
		{name: "return 0", args: args{haystack: "sadbutsad", needle: "sad"}, want: 0},
		{name: "-1", args: args{haystack: "leetcode", needle: "leeto"}, want: -1},
		{name: "3", args: args{haystack: "aabaabaafa", needle: "aabaaf"}, want: 3},
		{name: "6", args: args{haystack: "ababcaababcaabc", needle: "ababcaabc"}, want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strStr()=%v, want: %v", got, tt.want)
			}
		})
	}
}
