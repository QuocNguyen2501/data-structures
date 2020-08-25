package main

import "fmt"

func main(){
	tree := &BinaryTree{}
	tree.Insert(40).
		Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60).
		Insert(50).
		Insert(60).
		Insert(55).
		Insert(85).
		Insert(15).
		Insert(5).
		Insert(-10)

	tree.node.PrintInOrder()
	fmt.Println("============================")

	tree.Delete(-20)
	tree.Delete(15)
	tree.Delete(5)
	tree.Delete(-10)

	tree.node.PrintInOrder()
}
