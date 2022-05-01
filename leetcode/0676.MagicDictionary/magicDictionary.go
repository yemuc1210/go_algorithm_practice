package lt676
//剑指offer专项突击版 64
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


