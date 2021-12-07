package offerS64
/*
剑指 Offer II 064. 神奇的字典
设计一个使用单词列表进行初始化的数据结构，单词列表中的单词 互不相同 。 如果给出一个单词，请判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单词存在于已构建的神奇字典中。

实现 MagicDictionary 类：

MagicDictionary() 初始化对象
void buildDict(String[] dictionary) 使用字符串数组 dictionary 设定该数据结构，dictionary 中的字符串互不相同
bool search(String searchWord) 给定一个字符串 searchWord ，判定能否只将字符串中 一个 字母换成另一个字母，使得所形成的新字符串能够与字典中的任一字符串匹配。如果可以，返回 true ；否则，返回 false 。/*


*/
type MagicDictionary struct {
	dictionary []string
}


/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{
		dictionary: []string{},
	}
}


func (this *MagicDictionary) BuildDict(dictionary []string)  {
	this.dictionary = dictionary
}

// 其实就是比较两个字符串，直接暴力就好
func (this *MagicDictionary) Search(searchWord string) bool {
	n := len(this.dictionary)
	ans :=0
	for i :=0; i < n; i++ {
		ans =0
		for j :=0; j<len(this.dictionary[i]); j++ {
			if len(this.dictionary[i]) == len(searchWord) && this.dictionary[i][j] != searchWord[j]{
				ans++
			}
		}
		if ans == 1{
			return true
			break
		}
	}

	return false
}



/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
