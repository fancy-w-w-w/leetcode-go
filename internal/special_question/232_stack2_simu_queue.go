package specialquestion

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func ConstructorMyQueue() MyQueue {
	return MyQueue{
		stack1: []int{},
		stack2: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stack2) == 0 {
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2, this.stack1 = append(this.stack2, this.stack1[i]), this.stack1[:len(this.stack1)-1]
		}
	}
	res := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return res
}

func (this *MyQueue) Peek() int {
	if len(this.stack2) == 0 {
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2, this.stack1 = append(this.stack2, this.stack1[i]), this.stack1[:len(this.stack1)-1]
		}
	}
	return this.stack2[len(this.stack2)-1]

}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
