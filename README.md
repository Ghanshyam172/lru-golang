# lru-golang
Simple LRU implementation using your own doubly linked list in Go language.

This LRU Implementation utilizes following data structures:-
 1) Doubly Linked List :- To make delete operartions efficient and to maintain order such that new elements are added from Front and           deleted from Rear end, which takes O(1) time. Elements at the Front are most recently used elements and at the Rear are least recently     used.
 2) Map : - To make get(key) and set(key, value) operations efficient. Takes O(1) to insert and retrieve an element.
 
 Note:-  Feedback for improvements is always welcomed.
