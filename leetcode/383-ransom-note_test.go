package leetcode

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "return false 1", args: args{ransomNote: "a", magazine: "b"}, want: false},
		{name: "return false 2", args: args{ransomNote: "aa", magazine: "ab"}, want: false},
		{name: "return true 1", args: args{ransomNote: "aa", magazine: "aab"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ransomNote(b *testing.B)  {
	for i:=0; i<b.N;i++ {
		validAnagram("aa", "aab")
	}
}
