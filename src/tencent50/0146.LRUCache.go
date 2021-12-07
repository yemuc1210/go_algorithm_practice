package tencent50
/*中
146. LRU 缓存机制
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：
	LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
	int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
	void put(int key, int value) 如果关键字已经存在，则变更其数据值；
		如果关键字不存在，则插入该组「关键字-值」。
		当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，
			从而为新的数据值留出空间。
进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
*/
/*
	最久未使用，使用链表进行模拟 不过链表对于容量控制支持不好
	map : 可行
		如何获得最久未使用的？  
		答:map[int]*ListNode    将频次作为关键词 value是链表（结构参照 图的邻接表）

*/
type LinkedNode struct{
	key,val int
	prev,next *LinkedNode
}
func initLinkedNode(key,val int) *LinkedNode {
	return &LinkedNode{
		key: key,
		val: val,
	}
}

type LRUCache struct {
	size,capcity int
	cache map[int]*LinkedNode
	head,tail *LinkedNode
}


func Constructor146(capacity int) LRUCache {
	list := LRUCache{
		size: 0,
		capcity: capacity,
		cache: map[int]*LinkedNode{},
		head: initLinkedNode(0,0),
		tail: initLinkedNode(0,0),
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

func (this *LRUCache) Get(key int) int {
	if _,ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.val
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok {
        node := initLinkedNode(key, value)
        this.cache[key] = node
        this.addToHead(node)
        this.size++
        if this.size > this.capcity{
            removed := this.removeTail()
            delete(this.cache, removed.key)
            this.size--
        }
    } else {
        node := this.cache[key]
        node.val = value
        this.moveToHead(node)
    }
}


func (this *LRUCache) addToHead(node *LinkedNode) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) removeNode(node *LinkedNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *LinkedNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) removeTail() *LinkedNode {
    node := this.tail.prev
    this.removeNode(node)
    return node
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */