package offerS28

//lt430
  type Node struct {
      Val int
      Prev *Node
      Next *Node
      Child *Node
  }
 

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
