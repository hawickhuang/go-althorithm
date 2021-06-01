package mypprof

import (
	"math/rand"
	"time"
)

func GetRandInt(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(10000000))
	}
	return nums
}

func BubbleSort(arr []int) []int {
	l := len(arr)

	if l <= 1 {
		return arr
	}

	for i := 0; i < l-1; i++ {
		for j := l-1; j > i; j-- {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}