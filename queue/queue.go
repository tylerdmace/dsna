package queue

import "github.com/tylerdmace/dsna/stack"

type Queue struct {
	inStack stack.Stack
	outStack stack.Stack
}

func (q *Queue) Enqueue(i int) {
	q.inStack.Push(i)
}

func (q *Queue) Dequeue() (int, error) {
	q.shiftStack()
	return q.outStack.Pop()
}

func (q *Queue) shiftStack() {
	if q.outStack.IsEmpty() {
		s := q.inStack.Size()
		for i := 0; i < s; i++ {
			e, _ := q.inStack.Pop()
			q.outStack.Push(e)
		}
	}
}
