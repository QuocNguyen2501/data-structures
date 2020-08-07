# Linked List
**_Linked List_** is an **_abstract data structure_** and provides similar to an array, but with the big advantage that inserting or deleting an element in the list is very cheap, compared to doing so in an array,
where we need to shift all elements after the current position.

In while arrays keep all the elements in the same block of memory, next to each other, 
Linked List contains all elements scattered around memory, by storing a pointer to the next element.

However, Linked List also has a disadvantage over arrays that is when we want to pick an element
in the middle of the list, we cannot access to that element by index directly like array.
we don't know its address, so we need to start at the beginning of the list, and iterate on the list until we find it.


## 3 kinds of linked list
_**Single linked list**_
    ![single-linked-list](https://user-images.githubusercontent.com/5207501/89659969-daa6a780-d8fa-11ea-8a44-501d8062c4de.PNG)

_**Double linked list**_
    ![double-linked-list](https://user-images.githubusercontent.com/5207501/89659974-dd090180-d8fa-11ea-87eb-b82ac339779d.PNG)
    
_**Circular linked lis**t_
    ![circular-linked-list](https://user-images.githubusercontent.com/5207501/89659973-dc706b00-d8fa-11ea-9eed-138c64db486d.PNG)