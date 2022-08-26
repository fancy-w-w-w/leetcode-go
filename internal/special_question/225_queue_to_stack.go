package specialquestion

// 队列实现栈
type MyStack struct {
	queue []int
	top   int
}

func ConstructorMyStack() MyStack {
	return MyStack{
		queue: []int{},
		top:   -1,
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	this.top++
}

func (this *MyStack) Pop() int {
	top := this.Top()
	this.queue = this.queue[:this.top]
	this.top--
	return top
}

func (this *MyStack) Top() int {
	return this.queue[this.top]
}

func (this *MyStack) Empty() bool {
	return this.top < 0
}
