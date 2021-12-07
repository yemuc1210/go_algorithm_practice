package lt430


type Node struct {
      Val int
      Prev *Node
      Next *Node
      Child *Node
 }

/*
430. 扁平化多级双向链表
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。


多链表变二叉树，先序输出ojbk   dfs

遍历某个节点node,如果child不为空，则将child链表扁平化，并插入当前node和下一个node之间  dfs过程
为了插入，需要知道扁平化的链表的最后一个节点last：三步走
 	（1）node 和node 下一个节点next断开
	 （2）将node和child相连
	 （3）将last和next相连  
	 这就串起来了
*/
func flatten(root *Node) *Node {
	dfs(root)
	return root
}

func dfs(node *Node)(last *Node) {
	cur := node
	for cur != nil {
		//找到next
		next := cur.Next

		//node child处理
		if cur.Child != nil {
			//递归处理
			childList := dfs(cur.Child)

			next = cur.Next   //暂存
			//node 和child相连
			cur.Next = cur.Child   //之前保存cur.Next，不然就丢了
			cur.Child.Prev = cur 

			//第三步，last和next链接
			if next != nil {
				childList.Next = next
				next.Prev = childList
			}


			//child置空，避免dfs中再次遍历
			cur.Child = nil
			last = childList   //返回last
		}else{
			last = cur
		}

		cur = next
	}
	return 

}