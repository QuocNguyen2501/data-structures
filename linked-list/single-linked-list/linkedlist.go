package main

import (
	"fmt"
	"sync"
)

type Item interface {}

type Node struct {
	content string
	next *Node
}

type ItemLinkedList struct{
	head *Node
	size int
	lock sync.RWMutex
}

func (ll *ItemLinkedList) Append(t string)  {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	appendNode := Node{t, nil}

	if ll.head == nil {
		ll.head = &appendNode
	}else {
		last := ll.head
		for last.next != nil {
			last = last.next
		}
		last.next = &appendNode
	}
	ll.size++
}


func (ll *ItemLinkedList) Insert(i int, t string) error{
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if i < 0|| i> ll.size{
		return fmt.Errorf("Index out of bound")
	}
	addNode := Node{t, nil}
	if i == 0 {
		addNode.next = ll.head
		ll.head = &addNode
	}
	node := ll.head
	for j:=0; j< i-2; j++ {
		node = node.next
	}
	addNode.next = node.next
	node.next = &addNode
	ll.size++
	return nil
}

func (ll *ItemLinkedList) RemoveAt(i int) error{
	ll.lock.Lock()
	defer ll.lock.Unlock()
	if i < 0 || i> ll.size {
		return fmt.Errorf("Index out of bound")
	}
	node:= ll.head
	for j:= 0; j< i-1;j++{
		node = node.next
	}
	removeNode := node.next
	node.next = removeNode.next
	ll.size--
	return nil
}

func (ll *ItemLinkedList) IndexOf(t string) int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()

	j:= 0
	node:= ll.head
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
