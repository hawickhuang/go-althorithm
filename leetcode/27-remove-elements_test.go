package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T)  {
	type args struct {
		nums []int
		val int
	}

	tests := []struct{
		name string
		args args
		want int
	}{
		{name: "return 2", args: args{nums: []int{3,2,2,3}, val: 3}, want: 2},
		{name: "return 5", args: args{nums: []int{0,1,2,2,3,0,4,2}, val: 2}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.nums, tt.args.val)	; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements()=%v, want: %v", got, tt.want)
			}
		})
	}
}
