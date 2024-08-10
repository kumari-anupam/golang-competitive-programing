package main

import "fmt"

type LRUNode struct {
	prev *LRUNode
	next *LRUNode
	key  int
	val  int
}

type LRUCache struct {
	head     *LRUNode
	tail     *LRUNode
	lruMap   map[int]*LRUNode
	capacity int
}

func constructor(capacity int) LRUCache {
	cache := LRUCache{
		head:     &LRUNode{key: 0, val: 0},
		tail:     &LRUNode{key: 0, val: 0},
		lruMap:   make(map[int]*LRUNode),
		capacity: capacity,
	}

	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (l *LRUCache) insertAfterHead(node *LRUNode) {
	curAfterHead := l.head.next
	node.next = curAfterHead
	node.prev = l.head
	l.head.next = node
	curAfterHead.prev = node
}

func deleteNode(node *LRUNode) {
	nextNode := node.next
	prevNode := node.prev
	prevNode.next = node.next
	nextNode.prev = node.prev
}

func (l *LRUCache) get(key int) int {
	if val, ok := l.lruMap[key]; ok {
		lNode := val
		deleteNode(lNode)
		l.insertAfterHead(lNode)
		return lNode.val
	}

	return -1
}

func (l *LRUCache) put(key, value int) {
	if val, ok := l.lruMap[key]; ok {
		lnode := val
		lnode.val = value
		deleteNode(lnode)
		l.insertAfterHead(lnode)
	} else {
		if len(l.lruMap) == l.capacity {
			lnode := l.tail.prev
			delete(l.lruMap, lnode.val)
			deleteNode(lnode)
		}
		node1 := &LRUNode{
			key: key,
			val: value,
		}
		l.insertAfterHead(node1)
		l.lruMap[key] = node1
	}
}
func main() {
	lru := constructor(2)
	lru.put(1, 1)
	lru.put(2, 2)
	fmt.Println(lru.get(1))
	lru.put(3, 3)
	fmt.Println(lru.get(2))
	lru.put(4, 4)
	fmt.Println(lru.get(1))
	fmt.Println(lru.get(3))
	fmt.Println(lru.get(4))
}
