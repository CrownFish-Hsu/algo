package main

type MyCircularQueue struct {
	queue []int
	l     int
	r     int
	size  int
	limit int
}

// 同时在队列里的元素个数 最大K个元素
func Constructor(k int) MyCircularQueue {
	queue := make([]int, k)
	return MyCircularQueue{queue, 0, 0, 0, k}
}

// 如果队列满了，false
// 如果没满 加入v
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	} else {
		this.queue[this.r] = value
		if this.r == this.limit-1 {
			this.r = 0
		} else {
			this.r++
		}

		this.size++
		return true
	}
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	} else {
		if this.l == this.limit-1 {
			this.l = 0
		} else {
			this.l++
		}

		this.size--
		return true
	}
}

// 弹出队列头部数字
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.l]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	last := this.r - 1
	if last == -1 {
		last = this.limit - 1
	}

	return this.queue[last]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.limit
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
