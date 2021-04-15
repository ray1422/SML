package utils

import (
	"fmt"
	"sync"
)

type Item interface{}

// Stack the stack of Items
type Stack struct {
	items []Item
	lock  sync.RWMutex
}

// New creates a new ItemStack
func NewStack() *Stack {
	s := &Stack{}
	s.items = []Item{}
	return s
}

// Print prints all the elements
func (s *Stack) Print() {
	fmt.Println(s.items)
}

// Push adds an Item to the top of the stack
func (s *Stack) Push(t Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

// Pop removes an Item from the top of the stack
func (s *Stack) Pop() Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}
func (s *Stack) Top() Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	return item
}
