package main

import (
	"container/list"
)

// 请你设计一个用于存储字符串计数的数据结构，并能够返回计数最小和最大的字符串。
//  双向链表+哈希
type node struct {
    keys  map[string]struct{}
    count int
}

type AllOne struct {
    *list.List
    nodes map[string]*list.Element
}

func Constructor() AllOne {
    return AllOne{list.New(), map[string]*list.Element{}}
}

func (l *AllOne) Inc(key string) {
    if cur := l.nodes[key]; cur != nil {
        curNode := cur.Value.(node)
        if nxt := cur.Next(); nxt == nil || nxt.Value.(node).count > curNode.count+1 {
            l.nodes[key] = l.InsertAfter(node{map[string]struct{}{key: {}}, curNode.count + 1}, cur)
        } else {
            nxt.Value.(node).keys[key] = struct{}{}
            l.nodes[key] = nxt
        }
        delete(curNode.keys, key)
        if len(curNode.keys) == 0 {
            l.Remove(cur)
        }
    } else { // key 不在链表中
        if l.Front() == nil || l.Front().Value.(node).count > 1 {
            l.nodes[key] = l.PushFront(node{map[string]struct{}{key: {}}, 1})
        } else {
            l.Front().Value.(node).keys[key] = struct{}{}
            l.nodes[key] = l.Front()
        }
    }
}

func (l *AllOne) Dec(key string) {
    cur := l.nodes[key]
    curNode := cur.Value.(node)
    if curNode.count > 1 {
        if pre := cur.Prev(); pre == nil || pre.Value.(node).count < curNode.count-1 {
            l.nodes[key] = l.InsertBefore(node{map[string]struct{}{key: {}}, curNode.count - 1}, cur)
        } else {
            pre.Value.(node).keys[key] = struct{}{}
            l.nodes[key] = pre
        }
    } else { // key 仅出现一次，将其移出 nodes
        delete(l.nodes, key)
    }
    delete(curNode.keys, key)
    if len(curNode.keys) == 0 {
        l.Remove(cur)
    }
}

func (l *AllOne) GetMaxKey() string {
    if b := l.Back(); b != nil {
        for key := range b.Value.(node).keys {
            return key
        }
    }
    return ""
}

func (l *AllOne) GetMinKey() string {
    if f := l.Front(); f != nil {
        for key := range f.Value.(node).keys {
            return key
        }
    }
    return ""
}
//  

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */