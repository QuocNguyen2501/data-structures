package main

import "testing"

var tree BinaryTree

func Test_insert(t *testing.T){
	tree.Insert(100)
	if tree.node.data != 100 {
		t.Errorf("wrong root, expected root is 100 and get %d",tree.node.data)
	}
}

func Test_insert_node(t *testing.T){
	tree.Insert(50).Insert(120).Insert(30).Insert(60)
	if tree.node.left.data != 50 {
		t.Errorf("wrong left node, expected left node is 50 and get %d",tree.node.left.data)
	}
	if tree.node.right.data != 120 {
		t.Errorf("wrong right node, expected right node is 120 and get %d",tree.node.right.data)
	}
}

func Test_find_item(t *testing.T){
	n, err := tree.FindItem(30)
	if err != nil {
		t.Errorf("Wrong result, expect node is 30")
	}
	if n.data != 30 {
		t.Errorf("Wrong result, expect node is 30 and get %d",n.data)
	}
}

func Test_remove_item(t *testing.T){
	tree.Delete(30)
	n, err := tree.FindItem(30)
	if n!= nil || err == nil {
		t.Errorf("Expect node is removed")
	}
}
