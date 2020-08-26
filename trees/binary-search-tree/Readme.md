# Binary Search Tree
It is a binary tree extension with several optional restrictions. 
The left child value of node be less than the parent's value, and 
the right child value of node always be greater than the parent's value.


![1_h8w4gpvpx5_n1kvhzWE0Fw](https://user-images.githubusercontent.com/5207501/90334822-8be3c680-dffa-11ea-95a4-31ac2e191936.png)

It's very useful for search operations since we can accurately determine at each node whether
the value is in the left or right sub-tree.

## Searching
We look at the above picture, we need search 1 in the tree: 
- we begin with the first node (root node). 
1 is less than 10, we go to the left side and ignore all the right nodes of root node. 

- Similarly, we continue with the next node (the node with value is 5), 1 is less than 5, we go to the left side.

- Go to the left side we get the value of next node is 1, 1 equal to 1. BINGO!!! We found the node that we need.
We don't need to traversal all the value like array.

