package structs

import "sort"

// 学习MySQL索引时，对B B+ 进行实践
// 目前go 版本是1.16.5  不支持泛型
type DateType int

// 多路平衡查找树
// 阶数：一个节点最多有多少个孩子节点   字母m
// 定义：
// 1. 每个结点最多有m-1个关键字
// 2. 根结点最少可以只有1个关键字
// 3. 非根结点至少有Math.ceil(m/2)-1个关键字。
// 4. 每个结点中的关键字都按照从小到大的顺序排列，
// 	每个关键字的左子树中的所有关键字都小于它，而右子树中的所有关键字都大于它。
// 5. 所有叶子结点都位于同一层，或者说根结点到每个叶子结点的长度都相同
type BTreeNode struct {
	// 数据存储使用切片? map
	// 需要顺序，使用切片
	KeyWords []int
	// 子树
	//  * 1 * 2 * 3 *  --> KeyWords有3个，而Child应有4个
	Child []*BTreeNode
	// 父
	Parent *BTreeNode
	// 节点关键字个数
	keyNum int
	// 是否为根节点
	ifRoot bool
}

// 定义树
type BTree struct {
	Root *BTreeNode
	// 阶数：每一个节点 最多有多少个子节点
	M int
}

func getMinKeyWords(M int) int {
	return M/2 - 1
}

// 构造器
func ConstructorOfBTree(M int) *BTree {
	return &BTree{
		M: M,
		Root: &BTreeNode{
			ifRoot: true,
		},
	}
}

// search：关键字val
// 返回节点，以及在节点的位置
// 若不存在，则返回的是最终的叶子节点的位置，也就是待插入的位置
func (bt *BTree) Search(val int) (bool, int, *BTreeNode) {
	node := &BTreeNode{}
	var idx int
	if bt == nil {
		return false, 0, nil
	}
	curNode := bt.Root
	for curNode != nil {
		// 节点内部进行二分查找
		// 官方库 sort.Search()  返回第一个为true的值
		// slice根据升序排列
		idx = sort.Search(curNode.keyNum, func(i int) bool {
			return curNode.KeyWords[i] >= val
		})
		if curNode.KeyWords[idx] == val {
			return true, idx, curNode
		}
		// 否则，需要子树查询
		//  * 1 * 2 * 4 *  --> KeyWords有3个，而Child应有4个
		// 若查找3 ，此时idx是2 ，所以应该查Child[2]，因为Child[3]里面都比4大
		if curNode.Child[idx] != nil {
			node = curNode
		} // 保存
		curNode = curNode.Child[idx]
	}
	return false, idx, node
}

// 分裂操作
// 当关键字个数超出限制时，需要分裂
func (bt *BTreeNode) Split(M int) *BTreeNode {
	// 具体来说，是某个节点需要分裂
	// 节点的关键字限制：
	// root: 1<=  num  <= m-1
	// normal: m/2-1  <= num <= m-1

	newNode := &BTreeNode{}
	parent := bt.Parent

	if parent == nil {
		parent = &BTreeNode{}
	}
	// 分裂思路：
	// 以节点中间元素为分界。若是偶数个数，则中间两个数任意选
	// 中间节点插入到父节点，重复上述操作，知道所有节点符合B数的规则
	mid := bt.keyNum/2 
	midVal := bt.KeyWords[mid]
	// 右半部分 分裂生成新节点    例：M=4   keyNum=4  mid=2  
	// keyWords[mid]上升
	// 左半部分： keyWords[:mid]  Child[:mid+1]
	// 右半部分: keyWord[mid+1:]  Child[mid+1: ]
	newNode.keyNum = M - mid
	bt.keyNum = mid - 1   //分裂后的左半部分   原先节点
	
	// 新节点
	copy(newNode.KeyWords, bt.KeyWords[mid+1:])
	// child
	copy(newNode.Child, bt.Child[mid+1:])
	// 左
	bt.KeyWords = bt.KeyWords[:mid]
	bt.Child = bt.Child[:mid+1]
	// 中间节点上升
	newNode.Parent = parent
	bt.Parent = parent
	// 中间节点插入父节点
	parentNum := parent.keyNum
	// 寻找合适的位置插入
	idx := sort.Search(parentNum, func(i int) bool {
		return parent.KeyWords[i] >= midVal 
	})
	// 在idx处插入  
	// 插入值  
	parent.KeyWords = append(parent.KeyWords[:idx],append([]int{midVal}, parent.KeyWords[idx:]...)... )
	parent.keyNum ++
	// 更新Child关系
	parent.Child = append(parent.Child[:idx],append([]*BTreeNode{bt,newNode}, parent.Child[idx:]...)... )
	
	// 若有必要，向上递归分裂
	if parent.keyNum >= M {
		return parent.Split(M)
	}
	return parent
}

// 插入一个值
// 1、先查看树中是否有此关键字
// 2、通过第1步也能确定会插入到哪个节点中
// 2、查看插入节点后是否需要分裂节点
func (bt *BTree) Insert(val int)  {
	// 1. search
	ok,_,node := bt.Search(val)
	if !ok {
		// 若不存在，则考虑插入  
		// 在node中寻找合适位置插入
		idx := sort.Search(node.keyNum, func(i int) bool {
			return node.KeyWords[i] >= val
		})
		// 在idx处插入
		node.KeyWords = append(node.KeyWords[:idx], append([]int{val}, node.KeyWords[idx:]...)... )
		node.keyNum++
		// 3. 考虑是否分裂
		if node.keyNum < bt.M {
			return 
		}else{
			// 分裂
			// 从node处开始 
			parent := node.Split(bt.M)
			for parent.Parent != nil {
				parent = node.Parent
			}
			// root
			// return parent
		}
	}
}
// 删除某个关键字
func (bt *BTree) Delete(val int) *BTreeNode{
	ok,i,node := bt.Search(val)
	if ok {
		t := node.DeleteNode(val,i)
	}
	return t
}

func (t *BTreeNode) DeleteNode(val,i int) *BTreeNode {
	// 判断该节点是否是终端节点
	return t
}