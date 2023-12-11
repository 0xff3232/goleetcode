package circularQueue

//Time: O(1)
//Space: O(N)

type MyCircularQueue struct {
	queue []int
	count int
	size  int
	head  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k), // queue
		count: 0,              // number of current elements
		size:  k,              // queue size
		head:  0,              // head index
	}

}

func (this *MyCircularQueue) EnQueue(value int) bool {
    // full case
    if this.IsFull() {
        return false
    }

    this.queue[(this.head + this.count) % this.size] = value // set value in queue

    this.count ++
    return true 
}

func (this *MyCircularQueue) DeQueue() bool {
    // empty case
    if this.IsEmpty() {
        return false
    }

    // update the head index
    this.head = (this.head + 1) % this.size

    // decrease the numer of elements by 1
    this.count--
    return true
}

func (this *MyCircularQueue) Front() int {
     // empty case
    if this.IsEmpty() {
        return -1
    }

    // return the head element
    return this.queue[this.head] 
}

func (this *MyCircularQueue) Rear() int {
     // empty case
    if this.IsEmpty() {
        return -1
    }

    // return the tail element
    return this.queue[(this.head + this.count - 1) % this.size]
}

func (this *MyCircularQueue) IsFull() bool {
    return this.count == this.size
}

func (this *MyCircularQueue) IsEmpty() bool {
    // queue is empty
    return this.count == 0
}

// will print current queued items
func (this *MyCircularQueue) Print() []int {
    items := make([]int, this.count) // Create a slice to hold the current elements in queue
    for i := 0; i < this.count; i++ {
        items[i] = this.queue[(this.head+i)%this.size] // Use modulo for wrap around
    }
    return items // Return the slice containing current elements
}

