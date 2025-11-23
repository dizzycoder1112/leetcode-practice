package main

func main() {}

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	n := len(this.queue)
	this.queue = append(this.queue, x)
	for range n {
		front := this.queue[0]
		this.queue = this.queue[1:]
		this.queue = append(this.queue, front)
	}
}

func (this *MyStack) Pop() int {
	front := this.queue[0]
	this.queue = this.queue[1:]
	return front
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
