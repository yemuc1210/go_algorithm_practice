package lt211
type TrieNode struct {
    children [26]*TrieNode
    isEnd    bool
}

func (t *TrieNode) Insert(word string) {
    node := t
    for _, ch := range word {
        ch -= 'a'
        if node.children[ch] == nil {
            node.children[ch] = &TrieNode{}
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

//参考208  还未做  字典树
type WordDictionary struct {
	triRoot *TrieNode

}


func Constructor() WordDictionary {
	return WordDictionary{&TrieNode{}}
}


func (this *WordDictionary) AddWord(word string)  {
	this.triRoot.Insert(word)
}

/*
根节点开始搜索
对于当前字符是字母和点号的情况，分别按照如下方式处理：

如果当前字符是字母，则判断当前字符对应的子结点是否存在，如果子结点存在则移动到子结点，继续搜索下一个字符，如果子结点不存在则说明单词不存在，返回 \text{false}false；

如果当前字符是点号，由于点号可以表示任何字母，因此需要对当前结点的所有非空子结点继续搜索下一个字符。

重复上述步骤，直到返回 \text{false}false 或搜索完给定单词的最后一个字符。

如果搜索完给定的单词的最后一个字符，则当搜索到的最后一个结点的 \textit{isEnd}isEnd 为 \text{true}true 时，给定的单词存在。


*/
func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
    dfs = func(index int, node *TrieNode) bool {
        if index == len(word) {
            return node.isEnd
        }
        ch := word[index]
        if ch != '.' {
            child := node.children[ch-'a']
            if child != nil && dfs(index+1, child) {
                return true
            }
        } else {
            for i := range node.children {
                child := node.children[i]
                if child != nil && dfs(index+1, child) {
                    return true
                }
            }
        }
        return false
    }
    return dfs(0, this.triRoot)

}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */