package main

import "fmt"

func main(){
	var linkedList ItemLinkedList
	linkedList.Append("First")
	linkedList.Append("Two")
	linkedList.Append("Three")

	index:= linkedList.IndexOf("Two")
	fmt.Print(index)
}