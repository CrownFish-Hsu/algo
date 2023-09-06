package main

type MinStack struct {
	stack_data []int
	stack_min  []int
}

func ConstructorMinStack() MinStack {
	stack_data := make([]int, 0)
	stack_min := make([]int, 0)
	return MinStack{stack_data: stack_data, stack_min: stack_min}
}

func (this *MinStack) Push(val int) {
	stackLen := len(this.stack_data)
	this.stack_data = append(this.stack_data, val)
	if stackLen == 0 {
		this.stack_min = append(this.stack_min, val)
	} else {
		min := val
		curMin := this.GetMin()
		if curMin < val {
			min = curMin
		}
		this.stack_min = append(this.stack_min, min)
	}
}

func (this *MinStack) Pop() {
	stackLen := len(this.stack_data)
	if stackLen > 0 {
		this.stack_data = this.stack_data[:stackLen-1]
		this.stack_min = this.stack_min[:stackLen-1]
	}
}

func (this *MinStack) Top() int {
	stackLen := len(this.stack_data)
	if stackLen == 0 {
		return -1
	} else {
		return this.stack_data[stackLen-1]
	}
}

func (this *MinStack) GetMin() int {
	stackLen := len(this.stack_min)
	if stackLen == 0 {
		return -1
	} else {
		return this.stack_min[stackLen-1]
	}
}
