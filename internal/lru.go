package internal

// LRU
type LRUCache struct {
	capacity   int
	m          map[int]*Node
	head, tail *Node
}

type Node struct {
	key       int
	value     int
	pre, next *Node
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.moveToHead(v)
		return v.value
	}
	return -1
}

func (this *LRUCache) moveToHead(node *Node) {
	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) deleteNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) removeTail() int {
	node := this.tail.pre
	this.deleteNode(node)
	return node.key
}

func (this *LRUCache) addToHead(node *Node) {
	this.head.next.pre = node
	node.next = this.head.next
	node.pre = this.head
	this.head.next = node
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.value = value
		this.moveToHead(v)
		return
	}

	if this.capacity == len(this.m) {
		rmkey := this.removeTail()
		delete(this.m, rmkey)
	}

	newNode := &Node{key: key, value: value}
	this.addToHead(newNode)
	this.m[key] = newNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		m:        map[int]*Node{},
		head:     head,
		tail:     tail,
	}
}
