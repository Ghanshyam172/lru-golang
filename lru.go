package lruimpl

type Node struct {
	Prev, Next *Node
	Data interface{}
}

type LRU struct {
	Front, Rear *Node
	hashTable map[interface{}]*Node
	totalFrames int
	length int
}

type NewData struct {
	key, value interface{}
}


func CreateLRU(capacity int) LRU {
	return LRU{
		totalFrames: capacity,
		hashTable:make(map[interface{}]*Node, capacity),
	}
}

func (lru *LRU) Set(key, value interface{}) {
	if node, ok := lru.hashTable[key]; ok {
		//if key already exists then delete the node from dll, move the node to the front of dll
		node.Data = NewData{
			key:   key,
			value: value,
		}
		lru.DeleteNode(node)
		lru.AddToFront(node)
	}else {
		node := &Node{
			Data: NewData{
				key:   key,
				value: value,
			},
		}
		lru.hashTable[key] = node
		//when cache is full, delete last node and add new node to the front of dll
		if lru.totalFrames == lru.length  {
			delete(lru.hashTable, lru.Rear.Data.(NewData).key)
			lru.DeleteNode(lru.Rear)
			lru.AddToFront(node)
		}else {// when cache has space available
			lru.length++
			lru.AddToFront(node)
		}
	}
}

func (lru *LRU) Get(key interface{}) interface{} {
	if node, ok := lru.hashTable[key]; ok {
		lru.AddToFront(node)
		return node.Data
	}
	return nil
}

func(lru *LRU) AddToFront(node *Node) {
	node.Next = lru.Front
	node.Prev = nil
	if lru.Front != nil {
		lru.Front.Prev = node
	}
	if lru.Rear == nil {
		lru.Rear = lru.Front
	}
	lru.Front = node
}

func (lru *LRU) DeleteNode(node *Node) {
	if node.Prev == nil {//if it is first node
		if node.Next != nil {
			node.Next.Prev = nil
		}
		lru.Front = node.Next
		node.Next = nil
	}else {//any other node than first node
		node.Prev.Next = node.Next
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}
		node.Prev = nil
		node.Next = nil
	}
}