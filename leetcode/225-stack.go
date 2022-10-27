package leetcode

type MyStack struct {
	headQueue []int
	backQueue []int
}

func constructor() MyStack {
	return MyStack{
		headQueue: make([]int, 0),
		backQueue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	for len(this.headQueue) != 0 {
		this.backQueue = append(this.backQueue, this.headQueue[0])
		this.headQueue = this.headQueue[1:]
	}
	this.headQueue = append(this.headQueue, x)
	for len(this.backQueue) != 0 {
		this.headQueue = append(this.headQueue, this.backQueue[0])
		this.backQueue = this.backQueue[1:]
	}
}

func (this *MyStack) Pop() int {
	if len(this.headQueue) == 0 {
		return 0
	}
	val := this.headQueue[0]
	this.headQueue = this.headQueue[1:]
	return val
}

func (this *MyStack) Top() int {
	val := this.Pop()
	this.Push(val)
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.headQueue) == 0
}
