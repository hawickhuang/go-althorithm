package leetcode

import "testing"

func Test_reverseString(t *testing.T)  {
	type args struct {
		s []byte
	}

	tests := []struct{
		name string
		args args
	}{
		{name: "1", args: args{s: []byte("test")}},
		{name: "2", args: args{s: []byte("hello world")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.s)
			reverseString(tt.args.s)
			t.Log(tt.args.s)
		})
	}
}
