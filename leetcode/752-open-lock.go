package leetcode

func openLock(deadends []string, target string) int {
	start := "0000"
	if target == start {
		return 0
	}
	dead := map[string]struct{}{}
	for _, s := range deadends {
		dead[s] = struct{}{}
	}
	if _, ok := dead[start]; ok {
		return -1
	}

	get := func(status string)(ret []string) {
		s := []byte(status)
		for i, b := range s {
			s[i] = b -1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))
			s[i] = b+1
			if s[i] > '9' {
				s[i] = '0'
			}
			ret = append(ret, string(s))
			s[i] = b
		}
		return
	}

	type Pair struct {
		status string
		step int
	}

	q := []Pair{{start, 0}}
	seen := map[string]bool{start: true}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			_, ok := dead[nxt]
			if !seen[nxt] && !ok {
				if nxt == target {
					return p.step + 1
				}
				seen[nxt] = true
				q = append(q, Pair{nxt, p.step+1})
			}
		}
	}
	return -1
}
