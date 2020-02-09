package datastructure

import "fmt"

type MyCircularQueue struct {
	front int
	rear int
	length int
	queue []int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) *MyCircularQueue {
	return &MyCircularQueue{
		front:0,
		rear:0,
		length:k,
		queue: make([]int, k),
	}
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull(){
		return false
	}

	this.queue[this.rear] = value
	this.rear = (this.rear + 1) % this.length
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty(){
		return false
	}

	this.front = (this.front + 1) % this.length
	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	return this.queue[this.front]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	return this.queue[this.rear]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.rear + 1) % this.length == this.front
}

func (this *MyCircularQueue) Print() {
	i := this.front
	for (i % this.length) != this.rear {
		fmt.Printf("%v->", this.queue[i])
		i++
	}
	fmt.Println()
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

