package main

type MyQueue struct {
	first  stack
	second stack
}

type stack []int

func (s *stack) pop() int {
	if len(*s) == 0 {
		return -1
	}
	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}

func (s *stack) push(x int) {
	*s = append(*s, x)
}
func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func Constructor() MyQueue {
	return MyQueue{make([]int, 0), make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.first.push(x)
}

func (this *MyQueue) Pop() int {
	if len(this.second) == 0 {
		for len(this.first) != 0 {
			this.second.push(this.first.pop())
		}
	}
	return this.second.pop()
}

func (this *MyQueue) Peek() int {
	if len(this.second) == 0 {
		for len(this.first) != 0 {
			this.second.push(this.first.pop())
		}
	}
	return this.second.peek()
}

func (this *MyQueue) Empty() bool {
	if len(this.first) == 0 && len(this.second) == 0 {
		return true
	}
	return false
}
