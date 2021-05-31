package offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])

	r, c := 0, cols -1

	for r < rows && c >= 0 {
		num := matrix[r][c]
		if target == num {
			return true
		} else if num > target {
			c -=1
		} else {
			r += 1
		}
	}
	return false
}