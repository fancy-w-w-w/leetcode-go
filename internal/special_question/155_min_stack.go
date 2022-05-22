package specialquestion

import "math"

type MinStack struct {
	stackSlice []*elem
	topIndex   int
	minVal     int
}
type elem struct {
	val    int
	minVal int
}

func Constructor() MinStack {
	return MinStack{
		stackSlice: []*elem{},
		topIndex:   -1,
		minVal:     math.MaxInt,
	}
}

func (this *MinStack) Push(val int) {
	minVal := min(val, this.minVal)
	this.stackSlice = append(this.stackSlice, &elem{val: val, minVal: min(val, minVal)})
	this.minVal = minVal
	this.topIndex++
}

func (this *MinStack) Pop() {
	this.topIndex--
	if this.topIndex == -1 {
		this.minVal = math.MaxInt
	} else {
		this.minVal = this.stackSlice[this.topIndex].minVal
	}
	this.stackSlice = this.stackSlice[:len(this.stackSlice)-1]
}

func (this *MinStack) Top() int {
	return this.stackSlice[this.topIndex].val
}

func (this *MinStack) GetMin() int {
	return this.stackSlice[this.topIndex].minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
