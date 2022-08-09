package specialquestion

import "math"

// 155. 最小栈
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

// 实现 MinStack 类:
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

// 最小栈 实现方式2
// type MinStack struct {
//     stack []int
//     minStack []int
// }

// func Constructor() MinStack {
//     return MinStack{
//         stack: []int{},
//         minStack: []int{math.MaxInt64},
//     }
// }

// func (this *MinStack) Push(x int)  {
//     this.stack = append(this.stack, x)
//     top := this.minStack[len(this.minStack)-1]
//     this.minStack = append(this.minStack, min(x, top))
// }

// func (this *MinStack) Pop()  {
//     this.stack = this.stack[:len(this.stack)-1]
//     this.minStack = this.minStack[:len(this.minStack)-1]
// }

// func (this *MinStack) Top() int {
//     return this.stack[len(this.stack)-1]
// }

// func (this *MinStack) GetMin() int {
//     return this.minStack[len(this.minStack)-1]
// }

// func min(x, y int) int {
//     if x < y {
//         return x
//     }
//     return y
// }
