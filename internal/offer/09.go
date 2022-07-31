package offer

// 剑指 Offer 09. 用两个栈实现队列
type CQueue struct {
	stack1 []int
	stack2 []int
}

func ConstructorCQueue() CQueue {
	return CQueue{
		stack1: []int{},
		stack2: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return -1
	}
	if len(this.stack2) == 0 {

		for i := len(this.stack1) - 1; i >= 0; i-- {
			v := this.stack1[i]
			this.stack1 = this.stack1[:i]
			this.stack2 = append(this.stack2, v)
		}
	}
	v := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return v
}
