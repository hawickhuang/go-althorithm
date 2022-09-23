package leetcode

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T)  {
	type args struct {
		nums []int
	}

	tests := []struct{
		name string
		args args
		want [][]int
	}{
		{name: "return [[-1,-1,2],[-1,0,1]]", args: args{nums: []int{-1,0,1,2,-1,-4}}, want: [][]int{[]int{-1,-1,2}, []int{-1,0,1}}},
		{name: "return []", args: args{nums: []int{0,1,1}}, want: [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum()=%v, want: %v", got, tt.want)
			}
		})
	}
}
