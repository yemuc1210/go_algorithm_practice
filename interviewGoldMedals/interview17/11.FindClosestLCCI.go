package interview17

// 单词距离
// 给定任意两个不同的单词，找出文件中的这两个单词的最短距离---相隔单词数
// 若寻找过程重复多次，如何优化
// 示例：
// words = ["I","am","a","student","from","a","university","in","a","city"],
//  word1 = "a", word2 = "student"
// 输出：1

// 遍历一遍的复杂度是O(n) 
// 重复多次，可以建矩阵，不过复杂度较高
// [0] : [a,b,c] [d,e,f]    值距离words[0]距离为1的元素有a,b,c 为2的有d,e,f


func findClosest(words []string, word1,word2 string) int {
	res := len(words) + 1  // 极大值
	idx1,idx2 := -1,-1
	for i,word := range words {
		if word == word1 {
			idx1 = i
		}else if word == word2 {
			idx2 = i
		}
		// 判断
		if idx1 >=0 && idx2 >=0 {
			res = min(res, abs(idx1-idx2))
		}
	}
	return res
}

func min(a,b int) int {
	if a<b {return a}
	return b
}
func abs(a int) int {
	if a<0 {return -a}
	return a
}