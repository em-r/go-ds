package ds

import (
	"errors"
	"sync"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type SinglyLinkedList struct {
	Head  *Node
	Size  int
	mutex sync.Mutex
}

func (s *SinglyLinkedList) incrementSize() {
	s.Size += 1
}

func (s *SinglyLinkedList) decrementSize() {
	s.Size -= 1
}

func (s *SinglyLinkedList) isEmpty() bool {
	return s.Head == nil
}

func (s *SinglyLinkedList) Items() []interface{} {
	var items []interface{}

	for currentNode := s.Head; currentNode != nil; currentNode = currentNode.Next {
		items = append(items, currentNode.Value)
	}

	return items
}

func (s *SinglyLinkedList) Append(n *Node) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	defer s.incrementSize()
	if s.isEmpty() {
		s.Head = n
		return
	}

	currentNode := s.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = n
}

func (s *SinglyLinkedList) Delete(index int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if index >= s.Size {
		return errors.New("index out of range")
	}

	defer s.decrementSize()
	if index == 0 {
		s.Head = s.Head.Next
		return nil
	}

	currentNode := s.Head
	var previousNode *Node
	for i := 0; i < index; i++ {
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	previousNode.Next = currentNode.Next

	return nil
}

func (s *SinglyLinkedList) Remove(value interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.isEmpty() {
		return
	}

	for {
		if s.Head.Value == value {
			s.Head = s.Head.Next
			s.decrementSize()
		} else {
			break
		}
	}

	var previousNode *Node
	currentNode := s.Head

	for currentNode != nil {
		if currentNode.Value == value {
			previousNode.Next = currentNode.Next
			currentNode = currentNode.Next
			s.decrementSize()
		} else {
			previousNode = currentNode
			currentNode = currentNode.Next
		}
	}
}

func (s *SinglyLinkedList) Prepend(n *Node) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	defer s.incrementSize()
	if s.isEmpty() {
		s.Head = n
		return
	}

	n.Next = s.Head
	s.Head = n
}

func (s *SinglyLinkedList) Reverse() {
	if s.Size <= 1 {
		return
	}

	currentNode := s.Head
	var previousNode *Node
	var tempNode *Node

	for currentNode != nil {
		tempNode = currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = tempNode
	}
	s.Head = previousNode
}

func (s *SinglyLinkedList) Clear() {
	if s.isEmpty() {
		return
	}
	s.Head = nil
	s.Size = 0
}

func (s *SinglyLinkedList) Unique() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.Size <= 1 {
		return
	}
	visited := make(map[interface{}]bool)
	currentNode := s.Head
	var previousNode *Node

	for currentNode != nil {
		if _, ok := visited[currentNode.Value]; ok {
			previousNode.Next = currentNode.Next
			s.decrementSize()
		} else {
			visited[currentNode.Value] = true
			previousNode = currentNode
		}
		currentNode = previousNode.Next
	}
}

func (s *SinglyLinkedList) Swap(i, j int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if i >= s.Size || j >= s.Size {
		return
	}
	if i == j || s.Size <= 1 {
		return
	}

	currentNodeX := s.Head
	var previousNodeX *Node
	xIndex := 0

	for currentNodeX != nil && xIndex != i {
		previousNodeX = currentNodeX
		currentNodeX = currentNodeX.Next
		xIndex++
	}

	if currentNodeX == nil {
		return
	}

	currentNodeY := s.Head
	var previousNodeY *Node
	yIndex := 0

	for currentNodeY != nil && yIndex != j {
		previousNodeY = currentNodeY
		currentNodeY = currentNodeY.Next
		yIndex++
	}

	if currentNodeY == nil {
		return
	}

	currentNodeX.Next, currentNodeY.Next = currentNodeY.Next, currentNodeX.Next
	previousNodeX.Next = currentNodeY
	previousNodeY.Next = currentNodeX
}
