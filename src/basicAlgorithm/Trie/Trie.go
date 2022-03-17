package Trie

// 前缀树节点
type Trie struct{
	// 本节点值
	Val interface{}
	// 子节点列表 使用map判断某个值是否存在，方便查询
	// Children map[interface{}]bool
	// 使用数组
	Children [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _,ch := range word {
		ch -= 'a'   // 得到数字/下标
		if node.Children[ch] == nil {
			// 新建
			node.Children[ch] = &Trie{}
		}
		node = node.Children[ch]  // 移动
	}
	node.isEnd = true  // 存入完整单词
}

func (this *Trie) SearchPrefix(prefix string) *Trie{
	node := this
	for _,ch := range prefix {
		ch -= 'a'
		if node.Children[ch] == nil {
			return nil
		}
		node = node.Children[ch]
	}
	return node
}
// 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node!=nil && node.isEnd
}

// 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}