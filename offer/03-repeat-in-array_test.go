package offer

import "testing"

func Test_findRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{2, 3, 1, 0, 2, 5, 3}}, want: 2},
		{name: "test2", args: args{nums: []int{1, 1}}, want: 1},
		{name: "test2", args: args{nums: []int{3, 4, 5, 6, 7, 8, 9, 1, 1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRepeatNumber2(t *testing.T) {
	r := findRepeatNumber2([]int{2, 3, 1, 0, 2, 5, 3})
	if r != 2 {
		t.Errorf("findRepeatNumber() = %d, want: %d", r, 2)
	}
	r = findRepeatNumber2([]int{1, 1})
	if r != 1 {
		t.Errorf("findRepeatNumber() = %d, want: %d", r, 1)
	}
	r = findRepeatNumber2([]int{2, 3, 4, 5, 2, 1, 1})
	if r != 1 && r != 2{
		t.Errorf("findRepeatNumber() = %d, want: %d", r, 1)
	}
}

