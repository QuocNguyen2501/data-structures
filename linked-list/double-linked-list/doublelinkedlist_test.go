package main

import "testing"

var ll ItemLinkedList

func TestAppend(t *testing.T){
	if !ll.IsEmpty(){
		t.Errorf("ll should be emptu")
	}

	ll.Append("item 1")

	if ll.IsEmpty() {
		t.Errorf("ll should not be empty")
	}

	if size:=ll.size; size != 1{
		t.Errorf("wrong count, expected 1 and got %d",size)
	}

	ll.Append("item 2")
	ll.Append("item 3")

	if size := ll.size; size!= 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestInsert(t *testing.T){
	err := ll.Insert(1,"item 4")
	if err !=nil {
		t.Errorf("unexpected error %s", err)
	}

	node:= ll.head
	for j := 0; j < 1 ; j++ {
		node = node.next
	}
	if node.content != "item 4" {
		t.Errorf("Wrong insert, expect position of")
	}
}

func TestRemoveAt(t *testing.T){
	err := ll.RemoveAt(1)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	node:= ll.head
	for j := 0; j < 1 ; j++ {
		node = node.next
	}
	if node.content == "item 4" {
		t.Errorf("Wrong remove at, expect position of")
	}
	if ll.size == 4 {
		t.Errorf("wrong count, expected 4 and got %d", ll.size)
	}
}

func TestIndexOf(t *testing.T) {
	if i := ll.IndexOf("item 1"); i != 0 {
		t.Errorf("expected position 0 but got %d", i)
	}
	if i := ll.IndexOf("item 2"); i != 1 {
		t.Errorf("expected position 1 but got %d", i)
	}
	if i := ll.IndexOf("item 3"); i != 2 {
		t.Errorf("expected position 2 but got %d", i)
	}
	if i := ll.IndexOf("item 4"); i != -1 {
		t.Errorf("expected position -1 but got %d", i)
	}
}

func TestHead(t *testing.T) {
	h := ll.head
	if "item 1" != h.content {
		t.Errorf("Expected `item 1` but got %s", h.content)
	}
}