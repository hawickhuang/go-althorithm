package leetcode

import "testing"

func Test_validAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "return true", args: args{s: "ats", t: "sat"}, want: true},
		{name: "return false", args: args{s: "ats", t: "sbt"}, want: false},
		{name: "return true with repeat char", args: args{s: "aaatsaa", t: "saaaaat"}, want: true},
		{name: "return false with repeat char", args: args{s: "aatsaa", t: "saaaaat"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("validAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_validAnagram(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		validAnagram("banana", "nabana")
	}
}