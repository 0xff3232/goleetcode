# Stack

## 1. Queue First-in-first-out Data Structure

### Goals:

- Understand the *definition* of FIFO and queue
- Be able to *implement* a queue by yourself
- Be familiar with the *built-in* queue strucure
- Use queue to solve problems
 
Problems:

- [Design Circular Queue](https://leetcode.com/explore/learn/card/queue-stack/228/first-in-first-out-data-structure/1337/)
- [Moving Average from Data Stream](https://leetcode.com/explore/learn/card/queue-stack/228/first-in-first-out-data-structure/1368/)

---
### 1.1 First-in-first-out Data Structure
*To implement a queue, we may use a dynamic array and an index pointing to the head of the queue.*

Two operations should be supported in our queue:
- *Enqueue* - element is added to end of que. (insert)

- *Dequeue* - first element is removed from the que. (delete)

---
### 1.2 Circular Queue
- if tail and head are on the same index empty
- if head is at one index and tail is at other index == 0 means we have full array
- tail is our enqueue (head points to the first enqueue if arr is empty)
- dequeue is our head pointer, head will point to first element until it isdequeue then it will i++
