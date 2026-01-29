package lru_cache

import "container/list"

// 解题思路：
// 可以把 LRU 想象为一摞书，每次抽出来（阅读或做笔记）和新加入的书都在最顶上，超过最大数量时，把底部的书淘汰

// 哈希表 (O(1) 查找)              双向链表 (O(1) 移动/删除)
// ┌─────────────────┐
// │ cache map       │             head (虚拟节点)
// │                 │               ↓
// │ key=1 → *Node1 ─┼─────────→  [Node1: k=1, v=10]
// │                 │               ↕ (prev/next指针)
// │ key=3 → *Node3 ─┼─────────→  [Node3: k=3, v=30]
// │                 │               ↕
// │ key=2 → *Node2 ─┼─────────→  [Node2: k=2, v=20]
// │                 │               ↓
// └─────────────────┘             tail (虚拟节点)
// 以上为图示，节点在哈希表中是随机存放的，通过指针连接为双向链表

// 解法一: 使用标准库 container/list
// Time: O(1) for both Get and Put, Space: O(capacity)
type entry struct {
	key, value int
}

type LRUCache1 struct {
	capacity int
	cache    map[int]*list.Element // 哈希表：key -> 链表元素
	list     *list.List            // 双向链表（头部=最近使用，尾部=最久未使用）
}

func Constructor1(capacity int) LRUCache1 {
	return LRUCache1{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (lc *LRUCache1) Get(key int) int {
	if elem, exists := lc.cache[key]; exists {
		// 将元素移到链表头部（标记为最近使用）
		lc.list.MoveToFront(elem)
		return elem.Value.(*entry).value
	}
	return -1
}

func (lc *LRUCache1) Put(key int, value int) {
	if elem, exists := lc.cache[key]; exists {
		// 更新已存在的元素
		lc.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
	} else {
		// 创建新元素并添加到链表头部
		elem := lc.list.PushFront(&entry{key: key, value: value})
		lc.cache[key] = elem

		// 检查容量，必要时删除最久未使用的元素
		if lc.list.Len() > lc.capacity {
			// 移除链表尾部元素
			oldest := lc.list.Back()
			if oldest != nil {
				lc.list.Remove(oldest)
				delete(lc.cache, oldest.Value.(*entry).key)
			}
		}
	}
}

// 解法二: 双向链表 + 哈希表
// Time: O(1) for both Get and Put, Space: O(capacity)
// 双向链表节点
type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node // 哈希表：key -> 链表节点
	head     *Node         // 虚拟头节点（最近使用）
	tail     *Node         // 虚拟尾节点（最久未使用）
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lc *LRUCache) Get(key int) int {
	if node, exists := lc.cache[key]; exists {
		// 将节点移到头部（标记为最近使用）
		lc.moveToHead(node)
		return node.value
	}
	return -1
}

func (lc *LRUCache) Put(key int, value int) {
	if node, exists := lc.cache[key]; exists {
		// 更新已存在的节点
		node.value = value
		lc.moveToHead(node)
	} else {
		// 创建新节点
		newNode := &Node{key: key, value: value}
		lc.cache[key] = newNode
		lc.addToHead(newNode)

		// 检查容量，必要时删除最久未使用的节点
		if len(lc.cache) > lc.capacity {
			removed := lc.removeTail()
			delete(lc.cache, removed.key)
		}
	}
}

// 辅助方法：将节点添加到头部
func (lc *LRUCache) addToHead(node *Node) {
	node.prev = lc.head
	node.next = lc.head.next
	lc.head.next.prev = node
	lc.head.next = node
}

// 辅助方法：移除节点
func (lc *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 辅助方法：将节点移到头部
func (lc *LRUCache) moveToHead(node *Node) {
	lc.removeNode(node)
	lc.addToHead(node)
}

// 辅助方法：移除尾部节点
func (lc *LRUCache) removeTail() *Node {
	node := lc.tail.prev
	lc.removeNode(node)
	return node
}
