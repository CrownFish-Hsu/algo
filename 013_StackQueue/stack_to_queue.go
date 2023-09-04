package main

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func ConstructorMyQueue() MyQueue {
	stackIn := make([]int, 0)
	stackOut := make([]int, 0)
	return MyQueue{stackIn, stackOut}
}

/*
*
倒数据：
1.out栈为空才能倒
2.in 必须一次倒完
*/
func (this *MyQueue) inToOut() {
	if len(this.stackOut) == 0 {
		for len(this.stackIn) > 0 {
			this.stackOut = append(this.stackOut, this.stackIn[len(this.stackIn)-1])
			this.stackIn = this.stackIn[:len(this.stackIn)-1]
		}
		for item := range this.stackIn {
			this.stackOut = append(this.stackOut, item)
		}
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn = append(this.stackIn, x)
	this.inToOut()
}

func (this *MyQueue) Pop() int {
	this.inToOut()

	if len(this.stackOut) > 0 {
		v := this.stackOut[len(this.stackOut)-1]
		this.stackOut = this.stackOut[:len(this.stackOut)-1]
		return v
	} else {
		return -1
	}
}

func (this *MyQueue) Peek() int {
	this.inToOut()

	if len(this.stackOut) > 0 {
		return this.stackOut[len(this.stackOut)-1]
	} else {
		return -1
	}
}

func (this *MyQueue) Empty() bool {
	return len(this.stackIn) == 0 && len(this.stackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
