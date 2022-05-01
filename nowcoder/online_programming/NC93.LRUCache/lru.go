package nowcoder93

//设计LRU缓存   最近最少使用

/**
 * lru design   lt146
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
*/

type LRUCache struct{
    size,capacity int
    cache map[int]*DLinkedNode //记录元素是否已经记录
    head,tail *DLinkedNode //维护首尾
}
type DLinkedNode struct{
    key,val int
    pre,next *DLinkedNode  //循环链表
}
func initDLinkedNode(key,val int) *DLinkedNode{
    return &DLinkedNode{
        key: key,
        val: val,
    }
}
func ConstructLRUCache(capacity int) LRUCache{
    l := LRUCache{
        cache: map[int]*DLinkedNode{},
        head: initDLinkedNode(0,0),
        tail: initDLinkedNode(0,0),
        capacity: capacity,
    }
    l.head.next = l.tail  //循环链表
    l.tail.next = l.head
    return l
}
func (this *LRUCache) removeNode(node *DLinkedNode){
    node.pre.next = node.next
    node.next.pre = node.pre
}
func (this *LRUCache) addToHead(node *DLinkedNode) {
    node.pre = this.head
    node.next = this.head.next
    this.head.next.pre = node  //先【后】 后 【前】
    this.head.next = node
}
func (this *LRUCache) moveToHead(node *DLinkedNode){
    this.removeNode(node)
    this.addToHead(node)
}
func (this *LRUCache) removeTail()*DLinkedNode{
    node := this.tail.pre
    this.removeNode(node)
    return node
}
//获得key对应的val
func (this *LRUCache) get(key int) int{
    //首先看存不存在
    if _,exist := this.cache[key]; !exist{
        return -1  //不存在
    }
    node := this.cache[key]
    //调整cache结构，最近访问的放到【队头】
    this.moveToHead(node)
    return node.val
}
func (this *LRUCache) set(key,val int) {
    if _,exist := this.cache[key];!exist{
        //插入新节点
        node := initDLinkedNode(key,val)
        //插入头部
        this.cache[key] = node
        this.addToHead(node)
        this.size ++
        //容量限制
        if this.size > this.capacity {
            //移除
            removed := this.removeTail()
            delete(this.cache, removed.key)
            this.size--
        }
    }else {
        //已经存在该节点  调整到队头
        node := this.cache[key]
        node.val = val
        this.moveToHead(node)
    }
}

func LRU( operators [][]int ,  k int ) []int {
    // write code here
    //初始化缓存
    l := ConstructLRUCache(k)   
    res := []int{}
    for _,op := range operators{
        if op[0] == 1 {
            //set操作
            l.set(op[1],op[2])
        }else if op[0] == 2{
            //get
            res = append(res, l.get(op[1]))
        }
    }
    return res
}