package main

type MyStack struct {
	queue []int
}

func ConstructorMyStack() MyStack {
	queue := make([]int, 0)
	return MyStack{queue}
}

// 将队列中头部元素一个个移至尾部
func (this *MyStack) Push(x int) {
	size := len(this.queue)
	this.queue = append(this.queue, x)
	for size > 0 {
		this.queue = append(this.queue, this.queue[0])
		this.queue = this.queue[1:]
		size--
	}
}

func (this *MyStack) Pop() int {
	size := len(this.queue)
	x := -1
	if size > 0 {
		x = this.queue[0]
		this.queue = this.queue[1:]
	}

	return x
}

func (this *MyStack) Top() int {
	size := len(this.queue)
	x := -1
	if size > 0 {
		x = this.queue[0]
	}

	return x
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
