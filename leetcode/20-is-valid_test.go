package leetcode

import (
	"reflect"
	"testing"
)

func Test_isValid(t *testing.T)  {
	type args struct {
		s string
	}

	tests := []struct{
		name string
		args args
		want bool
	}{
		{name: "return true", args: args{s: "()"}, want: true},
		{name: "return true", args: args{s : "()[]{}"}, want: true},
		{name: "return false", args: args{s: "(]"}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isValid()=%v, want: %v", got, tt.want)
			}
		})
	}
}
