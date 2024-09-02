package ltmath

type LRUCache struct {
	Size     int
	Capacity int
	Cache    map[int]*DoubleListNode
	Head     *DoubleListNode
	Tail     *DoubleListNode
}

func Constructor(capacity int) LRUCache {

	lruIns := LRUCache{
		Cache:    map[int]*DoubleListNode{},
		Head:     &DoubleListNode{Key: 0, Value: 0},
		Tail:     &DoubleListNode{Key: 0, Value: 0},
		Capacity: capacity,
	}

	lruIns.Head.Next = lruIns.Tail
	lruIns.Tail.Next = lruIns.Head

	return lruIns
}

func (lru *LRUCache) Get(key int) int {

	if _, ok := lru.Cache[key]; !ok {
		return -1
	}

	node := lru.Cache[key]

	lru.moveToHead(node)

	return node.Value
}

func (lru *LRUCache) Put(key int, value int) {

	if _, ok := lru.Cache[key]; ok {
		node := lru.Cache[key]
		node.Value = value
		lru.moveToHead(node)
		return
	}

	node := &DoubleListNode{Key: key, Value: value}
	lru.Cache[key] = node
	lru.addToHead(node)
	lru.Size++

	if lru.Size > lru.Capacity {
		removeNode := lru.removeTail()
		delete(lru.Cache, removeNode.Key)
		lru.Size--
	}

}

func (lru *LRUCache) moveToHead(node *DoubleListNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeNode(node *DoubleListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (lru *LRUCache) addToHead(node *DoubleListNode) {
	node.Prev = lru.Head
	node.Next = lru.Head.Next
	lru.Head.Next.Prev = node
	lru.Head.Next = node
}

func (lru *LRUCache) removeTail() *DoubleListNode {
	node := lru.Tail.Prev
	lru.removeNode(node)
	return node
}
