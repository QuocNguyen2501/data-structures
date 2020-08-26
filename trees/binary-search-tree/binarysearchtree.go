package main

import (
	"errors"
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

func (t *BinaryTree) Insert(data int64) *BinaryTree{
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

func (t *BinaryTree) FindMin() int64{
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

func (t *BinaryTree) FindMax() int64{
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

func (n *BinaryNode) findMaxWithParent(parent *BinaryNode) (*BinaryNode,*BinaryNode){
	if n == nil {
		return nil, parent
	}
	if n.right == nil {
		return n,parent
	}
	return n.right.findMaxWithParent(n)
}

func (t *BinaryTree) FindItem(v int64) (*BinaryNode,error) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.node.findItem(v)
}

func (n *BinaryNode) findItem(v int64) (*BinaryNode,error){
	if n == nil {
		return nil, errors.New("Item not found")
	}

	if v < n.data {
		return n.left.findItem(v)
	} else if v > n.data {
		return n.right.findItem(v)
	}
	return n,nil
}

func (t *BinaryTree) PrintInOrder(){
	t.lock.RLock()
	defer t.lock.RUnlock()
	t.node.printInOrder()
}

func (n *BinaryNode) printInOrder(){
	if n==nil {
		return
	}

	n.left.printInOrder()
	fmt.Println(n.data)
	n.right.printInOrder()
}

func (t *BinaryTree) PrintPreOrder(){
	t.lock.RLock()
	defer t.lock.RUnlock()
	t.node.printPreOrder()
}

func (n *BinaryNode) printPreOrder(){
	if n == nil {
		return
	}
	fmt.Println(n.data)
	n.left.printPreOrder()
	n.right.printPreOrder()
}

func (t *BinaryTree) PrintPostOrder(){
	t.lock.RLock()
	defer t.lock.RUnlock()
	t.node.printPostOrder()
}

func (n *BinaryNode) printPostOrder(){
	if n == nil {
		return
	}
	n.left.printPostOrder()
	n.right.printPostOrder()
	fmt.Println(n.data)
}

func (n *BinaryNode) ReplaceNode(parent,replacement *BinaryNode) error {
	if n == nil {
		return errors.New("Cannot replace a node with value's nil")
	}

	if n == parent.left {
		parent.left = replacement
		return nil
	}
	parent.right = replacement
	return nil
}

func (t *BinaryTree) Delete(v int64) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.node.delete(v,t.node)
}

func (n *BinaryNode) delete(v int64, parent *BinaryNode) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree.")
	}

	if v < n.data {
		return n.left.delete(v,n)
	} else if v >  n.data {
		return n.right.delete(v,n)
	}

	if n.left == nil && n.right == nil {
		n.ReplaceNode(parent,nil)
		return nil
	}

	if n.left == nil {
		n.ReplaceNode(parent,n.right)
		return nil
	}

	if n.right == nil {
		n.ReplaceNode(parent,n.left)
		return nil
	}

	replacement, replParent := n.left.findMaxWithParent(n)
	n.data = replacement.data

	return replacement.delete(replacement.data,replParent)
}

