package main

import (
	"fmt"
	"sync"
)

type BinaryNode struct{
	left *BinaryNode
	right *BinaryNode
	data int64
}

type BinaryTree struct {
	node *BinaryNode
	lock sync.RWMutex
}

func (t *BinaryTree) insert(data int64) *BinaryTree{
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.node == nil {
		t.node = &BinaryNode{data: data}
	}else {
		t.node.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	}
	if n.data == data {
		return
	}
	if data < n.data {
		if n.left == nil {
			n.left = &BinaryNode{data:data}
		}else {
			n.left.insert(data)
		}
	}else if data > n.data{
		if n.right == nil {
			n.right = &BinaryNode{data:data}
		}else {
			n.right.insert(data)
		}
	}
}

func (t *BinaryTree) findMin() int64{
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.node.findMin()
}

func (n *BinaryNode) findMin() int64 {
	if n.left == nil {
		return n.data
	}
	return n.left.findMin()
}

func (t *BinaryTree) findMax() int64{
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.node.findMax()
}

func (n *BinaryNode) findMax() int64{
	if n.right == nil{
		return n.data
	}
	return n.right.findMax()
}

func (n *BinaryNode) PrintInOrder(){
	if n==nil {
		return
	}

	n.left.PrintInOrder()
	fmt.Println(n.data)
	n.right.PrintInOrder()
}


