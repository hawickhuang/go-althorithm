package leetcode

func peakIndexInMountainArray(arr []int) int  {
	left, right := 0, len(arr)-2
	ans := 0
	for left <= right {
		mid := (left+right)/2
		if arr[mid] > arr[mid+1] {
			ans = mid
			right = mid-1
		} else {
			left = mid+1
		}
	}
	return ans
}
