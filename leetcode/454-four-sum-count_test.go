package leetcode

import (
	"reflect"
	"testing"
)

func Test_fourSumCount(t *testing.T)  {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
	}

	tests:=[]struct{
		name string
		args args
		want int
	}{
		{name: "return 2", args: args{nums1: []int{1,2}, nums2: []int{-2,-1}, nums3: []int{-1,2}, nums4: []int{0,2}}, want: 2},
		{name: "return 1", args: args{nums1: []int{1}, nums2: []int{0}, nums3: []int{0}, nums4: []int{-1}}, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.nums1, tt.args.nums2, tt.args.nums3, tt.args.nums4); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSumCount()=%v, want:%v", got, tt.want)
			}
		})
	}
}
