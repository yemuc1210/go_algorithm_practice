package offerS62

/*
实现前缀树
每个节点包含以下字段：指向子节点的指针数组children；isEnd bool 是否为字符串结尾
插入：
    从根开始，对于当前字符对应的子节点；
    1.子节点存在，沿着指针移动到子节点，继续下一字符
    2.不存在，创建新节点，记录在children数组对应位置，移动，继续
查找前缀：
    根开始，查找当前字符对应的子节点

*/
type Trie struct{
    children [26]*Trie
    isEnd bool
}

func Constructor() Trie{
    return Trie{}
}

func (t *Trie)Insert(word string){
    node := t
    for _,ch := range word{
        ch -= 'a'  //转数字
        if node.children[ch] == nil {
            //创建
            node.children[ch] = &Trie{}
        }
        node = node.children[ch]  //移动
    }
    node.isEnd = true
}

func (t *Trie)SearchPrefix(prefix string) *Trie {
    node := t
    for _,ch := range prefix {
        ch -= 'a'
        if node.children[ch] == nil {
            return nil
        }
        node = node.children[ch]
    }
    return node
}

func (t *Trie) Search(word string) bool {
    node := t.SearchPrefix(word)
    return node !=nil && node.isEnd
}

func (t *Trie)StartWith(prefix string)bool{
    return t.SearchPrefix(prefix) != nil
}
