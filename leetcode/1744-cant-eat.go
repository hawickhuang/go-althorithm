package leetcode

func canEat(candiesCount []int, queries [][]int) []bool {
	if len(queries) == 0 {
		return nil
	}
	res := make([]bool, len(queries))
	if len(candiesCount) == 0 {
		return res
	}

	sum := make([]int, len(candiesCount))
	sum[0] = candiesCount[0]
	for i := 1; i < len(candiesCount); i++{
		sum[i] = sum[i-1] + candiesCount[i]
	}

	for i, q := range queries {
		earliest := q[1] + 1
		latest := (q[1] +1) * q[2]

		min := 1
		if q[0] > 0 {
			min = sum[q[0]-1]+1
		}
		max := sum[q[0]]
		res[i] = !(earliest > max || latest < min)
	}

	return res
}
