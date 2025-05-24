// File contains the implementation of the data structures
// that are required in competitive programming

package main

import (
	"panic"
)


type Stack struct {
   sz  int
   values []interface{}
}

func (st *Stack) Push () {

}

func (st *Stack) Pop() interface{} {
    if st.size == 0 {
       panic("Stack is empty..")
	}
}

func (st * Stack) Top() {

}
