package main

import "testing"

var tree BinaryTree

func Test_insert(t *testing.T){
	tree.insert(100)
	if tree.node.data != 100 {
		t.Errorf("wrong root, expected root is 100 and get %d",tree.node.data)
	}
}

func Test_insert_node(t *testing.T){
	tree.insert(50).insert(120)
	if tree.node.left.data != 50 {
		t.Errorf("wrong left node, expected left node is 50 and get %d",tree.node.left.data)
	}
	if tree.node.right.data != 120 {
		t.Errorf("wrong right node, expected right node is 120 and get %d",tree.node.right.data)
	}
}