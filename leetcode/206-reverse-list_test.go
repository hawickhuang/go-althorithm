package leetcode

import (
	"reflect"
	"testing"
)



func Test_reverseList(t *testing.T)  {
	type args struct {
		head *ListNode
	}

	build := func(vals []int) *ListNode{
		var ant *ListNode
		prev := ant
		for i, val := range vals {
			node := &ListNode{val, nil }
			if i == 0 {
				prev, ant = node, node
			} else {
				prev.Next = node
			}
		}
		return ant
	}

	tests := []struct{
		name string
		args args
		want *ListNode
	}{
		{name: "return [5,4,3,2,1]", args: args{head: build([]int{1,2,3,4,5})}, want: build([]int{5,4,3,2,1})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList()=%v, want: %v", got, tt.want)
			}
		})
	}
}
