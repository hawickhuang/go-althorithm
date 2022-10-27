package leetcode

type MyQueue struct {
	head []int
	back []int
}

func Constructor() MyQueue {
	return MyQueue{
		head: make([]int, 0),
		back: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	for len(this.back) != 0 {
		this.head = append(this.head, this.back[len(this.back)-1])
		this.back = this.back[:len(this.back)-1]
	}
	this.head = append(this.head, x)
}

func (this *MyQueue) Pop() int {
	for len(this.head) != 0 {
		this.back = append(this.back, this.head[len(this.head)-1])
		this.head = this.head[:len(this.head)-1]
	}

	if len(this.back) == 0 {
		return 0
	}

	val := this.back[len(this.back)-1]
	this.back= this.back[:len(this.back)-1]
	return val
}

func (this *MyQueue) Peek() int {
	for len(this.head) != 0 {
		this.back = append(this.back, this.head[len(this.head)-1])
		this.head = this.head[:len(this.head)-1]
	}

	if len(this.back) == 0 {
		return 0
	}

	return this.back[len(this.back)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.back) + len(this.head) == 0
}
