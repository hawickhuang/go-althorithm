package leetcode

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		// TODO: Add test cases.
		{name: "return 0, 6", args: args{s: "cbaebabacd", p:"abc"}, wantAns: []int{0,6}},
		{name: "return 0,1,2", args: args{s:"abab", p:"ab"}, wantAns: []int{0,1,2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findAnagrams() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Benchmark_findAnagrams(b *testing.B)  {
	for i:=0;i<b.N;i++ {
		findAnagrams("abab", "ab")
	}
}
