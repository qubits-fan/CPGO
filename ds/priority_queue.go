package main

// Heap based implementation of Priority Queue it will support arbitary any element
// to compare

import (
	"fmt"
)

// Comparator should be if func (a < b) return -1
// if a == b return 0
// if a > b return 1

// PriorityQueue does not allow you to create heap of desired size
type PriorityQueue[T any] struct {
	minHeap int
	len     int
	hp      []T
	cmp     func(x, y T) int
}

func (p *PriorityQueue[T]) Push(val T) {
	if p == nil {
		return
	}
	// Push the element in last
	p.hp = append(p.hp, val)
	p.len++

	child := p.len - 1
	for child > 0 {
		parent := (child - 1) / 2
		if (p.cmp(p.hp[child], p.hp[parent]) * p.minHeap) > 0 {
			p.hp[parent], p.hp[child] = p.hp[child], p.hp[parent]
		} else {
			break
		}
		child = parent
	}
}

func (p *PriorityQueue[T]) Pop() T {
	if p.len == 0 {
		panic("Priority queue is empty..")
	}

	val := p.hp[0]
	p.hp[0] = p.hp[p.len-1]
	p.len--

	parent := 0
	for parent < p.len {
		lchild := 2*parent + 1
		rchild := 2*parent + 2

		if lchild >= p.len {
			break
		} else if rchild >= p.len {
			p.hp[lchild], p.hp[parent] = p.hp[parent], p.hp[lchild]
			parent = lchild
		} else {
			if (p.cmp(p.hp[lchild], p.hp[rchild]) * p.minHeap) > 0 {
				// Swap will happen from left child
				p.hp[lchild], p.hp[parent] = p.hp[parent], p.hp[lchild]
				parent = lchild
			} else {
				// Swap can happen right side
				p.hp[parent], p.hp[rchild] = p.hp[rchild], p.hp[parent]
				parent = rchild
			}
		}
	}
	return val
}

func (p *PriorityQueue[T]) Top() T {
	if p.len == 0 {
		panic("Heap is empty...")
	}
	return p.hp[0]
}

func NewPriorityQueue[T any](cmp func(x, y T) int, mh bool) *PriorityQueue[T] {
	pq := PriorityQueue[T]{
		len: 0,
		cmp: cmp,
		hp:  []T{},
	}

	// If minheap is allowed just revert the comparator function behaviour
	if mh {
		pq.minHeap = -1
	} else {
		pq.minHeap = 1
	}
	return &pq
}

func (p *PriorityQueue[T]) PrintPQ() {
	fmt.Printf("The MinPQ[%d], length [%d]\n", p.minHeap, p.len)
	for i := 0; i < p.len; i++ {
		fmt.Printf("%v ", p.hp[i])
	}
	fmt.Println()
}
