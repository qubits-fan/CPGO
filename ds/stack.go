package main

type Stack[T any] struct {
	len int
	st  []T
}

func (st *Stack[T]) Push(ele T) {
	st.st = append(st.st, ele)
	st.len++
}

func (st *Stack[T]) Pop() T {
	if st.len == 0 {
		panic("Stack is empty..")
	}

	st.len--
	val := st.st[st.len]
	st.st = st.st[0:st.len]
	return val
}

func (st *Stack[T]) Top() T {
	if st.len == 0 {
		panic("Stack is empty...")
	}
	return st.st[st.len-1]
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}
