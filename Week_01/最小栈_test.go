package Week_01_test

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {

	obj := Constructor()
	obj.Push(1)
	obj.Push(-1)
	obj.Push(-2)
	obj.Push(4)
	fmt.Println(obj.GetMin())
}

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	var top int
	if len(this.stack) == 0 {
		top = x
	} else {
		top = this.minStack[len(this.minStack)-1]
	}
	this.minStack = append(this.minStack, min(top, x))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
