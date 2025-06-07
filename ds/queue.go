package main

type Queue[T any] struct {
	len int
	q   []T
}

func (q *Queue[T]) Push(ele T) {
	q.q = append(q.q, ele)
	q.len++
}

func (q *Queue[T]) Front() T {
	if q.len == 0 {
		panic("Queue is empty...")
	}
	return q.q[0]
}

func (q *Queue[T]) Pop() T {
	if q.len == 0 {
		panic("Queue is empty...")
	}
	val := q.q[0]
	q.q = q.q[1:q.len]
	q.len--
	return val
}

func (q *Queue[T]) PopBack() T {
	if q.len == 0 {
		panic("Queue is empty...")
	}
	q.len--
	val := q.q[q.len]
	q.q = q.q[0:q.len]
	return val
}

func NewQueue[T any]() *Queue[T] {
	return new(Queue[T])
}
