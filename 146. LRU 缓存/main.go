package main

type Node struct {
	Key       int
	Value     int
	Pre, Next *Node
}

type LRUCache struct {
	capacity   int
	m          map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		capacity: capacity,
		m:        map[int]*Node{},
		head:     head,
		tail:     tail,
	}
}

func (t *LRUCache) moveToHead(n *Node) {
	t.deleteNode(n)
	t.addHead(n)
}

func (t *LRUCache) deleteNode(n *Node) {
	n.Pre.Next = n.Next
	n.Next.Pre = n.Pre
}
func (t *LRUCache) addHead(n *Node) {
	t.head.Next.Pre = n
	n.Next = t.head.Next
	n.Pre = t.head
	t.head.Next = n
}
func (t *LRUCache) removeTail() int {
	node := t.tail.Pre
	t.deleteNode(node)
	return node.Key
}

func (t *LRUCache) Get(key int) int {
	if v, ok := t.m[key]; ok {
		t.moveToHead(v)
		return v.Value
	}
	return -1
}

func (t *LRUCache) Put(key int, value int) {
	if v, ok := t.m[key]; ok {
		v.Value = value
		t.moveToHead(v)
		return
	}

	if t.capacity == len(t.m) {
		rmKey := t.removeTail()
		delete(t.m, rmKey)
	}

	newNode := &Node{Key: key, Value: value}
	t.addHead(newNode)
	t.m[key] = newNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
