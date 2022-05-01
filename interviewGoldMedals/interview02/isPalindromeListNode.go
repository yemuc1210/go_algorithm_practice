package interview02
/*
面试题 02.06. 回文链表

编写一个函数，检查输入的链表是否是回文的。

链表遍历，数据变数组，然后在数组中判断回文

你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

快慢指针，快指针走2，到达尾部时，慢指针处于中间  算，这也不能回溯或者记录访问数据，空间复杂度要求O(1)

利用递归本身的栈

快慢指针，修改链表后半部分，最后恢复结构
算法

整个流程可以分为以下五个步骤：

找到前半部分链表的尾节点。
反转后半部分链表。
判断是否回文。
恢复链表。
返回结果。
*/
func isPalindrome(head *ListNode) bool {
	root := head
	var flag bool = true
	var dfs func(node *ListNode) 
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}
		dfs(node.Next)
	
		if node.Val != root.Val{
			flag = false
		}
		root = root.Next
	}
	dfs(head)
	return flag
}

func reverseList(head *ListNode) *ListNode {
    var prev, cur *ListNode = nil, head
    for cur != nil {
        nextTmp := cur.Next
        cur.Next = prev
        prev = cur
        cur = nextTmp
    }
    return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
    fast := head
    slow := head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

func isPalindrome1(head *ListNode) bool {
    if head == nil {
        return true
    }

    // 找到前半部分链表的尾节点并反转后半部分链表
    firstHalfEnd := endOfFirstHalf(head)
    secondHalfStart := reverseList(firstHalfEnd.Next)

    // 判断是否回文
    p1 := head
    p2 := secondHalfStart
    result := true
    for result && p2 != nil {
        if p1.Val != p2.Val {
            result = false
        }
        p1 = p1.Next
        p2 = p2.Next
    }

    // 还原链表并返回结果
    firstHalfEnd.Next = reverseList(secondHalfStart)
    return result
}

