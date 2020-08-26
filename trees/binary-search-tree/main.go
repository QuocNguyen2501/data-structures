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

	fmt.Println("==========Print inorder============")
	tree.PrintInOrder()
	fmt.Println("============================")

	fmt.Println("==========Print preorder============")
	tree.PrintPreOrder()
	fmt.Println("============================")

	fmt.Println("==========Print post-order============")
	tree.PrintPostOrder()
	fmt.Println("============================")


	fmt.Println("==========Print tree after deleted nodes============")
	tree.Delete(-20)
	tree.Delete(15)
	tree.Delete(5)
	tree.Delete(-10)
	tree.PrintInOrder()
	fmt.Println("============================")
}
