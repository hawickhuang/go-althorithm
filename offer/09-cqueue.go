package offer

type CQueue struct {
	inStack  []int
	outStack []int
}

func NewCQueue() CQueue {
	return CQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (q *CQueue) AppendTail(value int) {
	q.inStack = append(q.inStack, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.outStack) == 0 && len(q.inStack) == 0 {
		return -1
	}
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			v := q.inStack[len(q.inStack)-1]
			q.inStack = q.inStack[:len(q.inStack)-1]
			q.outStack = append(q.outStack, v)

		}
	}

	v := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]

	return v
}
