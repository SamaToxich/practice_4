package main

type QueueNode[T any] struct {
	value T
	next  *QueueNode[T]
}

type Queue[T any] struct {
	head *QueueNode[T]
	tail *QueueNode[T]
}

func (q *Queue[T]) Enqueue(value T) {
	node := &QueueNode[T]{value: value}
	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.head == nil {
		var zero T
		return zero, false
	}
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return value, true
}

func (q *Queue[T]) Front() (T, bool) {
	if q.head == nil {
		var zero T
		return zero, false
	}
	return q.head.value, true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}
