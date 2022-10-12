package leetcode

func reverseWords(s string) string {
	bs := []byte(s)

	slow, fast := 0, 0
	for len(bs) > 0 && fast < len(bs) && bs[fast] == ' ' {
			fast++
	}

	for ; fast < len(bs); fast++ {
		if fast-1 >0 && bs[fast-1]==bs[fast] && bs[fast] == ' ' {
			continue
		}
		bs[slow] = bs[fast]
		slow++
	}

	if slow-1>0 && bs[slow-1] == ' ' {
		bs = bs[:slow-1]
	} else {
		bs = bs[:slow]
	}

	reverse(&bs, 0, len(bs)-1)

	i := 0
	for i<len(bs) {
		j := i
		for ; j < len(bs) && bs[j] != ' '; j++{}
		reverse(&bs, i, j-1)
		i = j
		i++
	}

	return string(bs)
}

func reverse(s *[]byte, start, end int)  {
	for start < end  {
		(*s)[start], (*s)[end] = (*s)[end], (*s)[start]
		start++
		end--
	}
}