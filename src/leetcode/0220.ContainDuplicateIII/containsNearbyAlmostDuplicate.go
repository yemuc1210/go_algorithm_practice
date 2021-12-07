package lt220

/*
剑指 Offer II 057. 值和下标之差都在给定的范围内
给你一个整数数组 nums 和两个整数 k 和 t 。
请你判断是否存在 两个不同下标 i 和 j，
使得 abs(nums[i] - nums[j]) <= t ，
同时又满足 abs(i - j) <= k 。

如果存在则返回 true，不存在返回 false。


滑动窗口+有序队列。每次用优先队列存放下窗口内的值（窗口大小为k），
这样在保证index差值小于k的情况下，
再用二分查找到有序队列里面找与nums[right]绝对差值最小的值，
检查是否小于等于t如果存在小于等于t的就返回True
*/
import "math/rand"

type node struct {
    ch       [2]*node
    priority int
    val      int
}

func (o *node) cmp(b int) int {
    switch {
    case b < o.val:
        return 0
    case b > o.val:
        return 1
    default:
        return -1
    }
}

func (o *node) rotate(d int) *node {
    x := o.ch[d^1]
    o.ch[d^1] = x.ch[d]
    x.ch[d] = o
    return x
}

type treap struct {
    root *node
}

func (t *treap) _put(o *node, val int) *node {
    if o == nil {
        return &node{priority: rand.Int(), val: val}
    }
    d := o.cmp(val)
    o.ch[d] = t._put(o.ch[d], val)
    if o.ch[d].priority > o.priority {
        o = o.rotate(d ^ 1)
    }
    return o
}

func (t *treap) put(val int) {
    t.root = t._put(t.root, val)
}

func (t *treap) _delete(o *node, val int) *node {
    if d := o.cmp(val); d >= 0 {
        o.ch[d] = t._delete(o.ch[d], val)
        return o
    }
    if o.ch[1] == nil {
        return o.ch[0]
    }
    if o.ch[0] == nil {
        return o.ch[1]
    }
    d := 0
    if o.ch[0].priority > o.ch[1].priority {
        d = 1
    }
    o = o.rotate(d)
    o.ch[d] = t._delete(o.ch[d], val)
    return o
}

func (t *treap) delete(val int) {
    t.root = t._delete(t.root, val)
}

func (t *treap) lowerBound(val int) (lb *node) {
    for o := t.root; o != nil; {
        switch c := o.cmp(val); {
        case c == 0:
            lb = o
            o = o.ch[0]
        case c > 0:
            o = o.ch[1]
        default:
            return o
        }
    }
    return
}

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
    set := &treap{}
    for i, v := range nums {
        if lb := set.lowerBound(v - t); lb != nil && lb.val <= v+t {
            return true
        }
        set.put(v)
        if i >= k {
            set.delete(nums[i-k])
        }
    }
    return false
}


