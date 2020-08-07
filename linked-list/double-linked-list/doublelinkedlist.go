package main

import (
	"fmt"
	"sync"
)

type Item interface {}

type Node struct {
	previous *Node
	content Item
	next *Node
}

type ItemLinkedList struct{
	head *Node
	size int
	lock sync.RWMutex
}

func (ll * ItemLinkedList) Append(t string){
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if ll.head == nil {
		ll.head = &Node{nil,t,nil}
	}else {
		last := ll.head
		for last.next != nil {
			last = last.next
		}
		last.next = &Node{last,t,nil}
	}
	ll.size++
}

func (ll *ItemLinkedList) Insert(i int, t string) error{
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if i < 0 || i > ll.size {
		return fmt.Errorf("Index out of bound")
	}

	if i == 0 {
		addNode := Node{nil,t, nil}
		addNode.next = ll.head
		ll.head = &addNode
	}
	node := ll.head
	for j:=0; j < i-2; j++ {
		node = node.next
	}
	addNode := Node{node,t, node.next}
	node.next = &addNode
	ll.size++
	return nil
}

func (ll *ItemLinkedList) RemoveAt(i int) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if i < 0 || i > ll.size {
		return fmt.Errorf("Index out of bound")
	}
	node := ll.head
	for j:=0; j< i-1;j++{
		node = node.next
	}
	removeNode := node.next
	node.next = removeNode.next
	node.next.previous = node
	ll.size--
	return nil
}

func (ll *ItemLinkedList) IndexOf(t string) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	j:=0
	node:=ll.head
	for {
		if node.content == t {
			return j
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		j++
	}
}

func (ll *ItemLinkedList) IsEmpty() bool {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	if ll.head == nil {
		return  true
	}
	return  false
}