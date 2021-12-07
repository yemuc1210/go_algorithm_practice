package lt820

/*
剑指 Offer II 065. 最短的单词编码
单词数组 words 的 有效编码 由任意助记字符串 s 和下标数组 indices 组成，
且满足：
	words.length == indices.length
	助记字符串 s 以 '#' 字符结尾
	对于每个下标 indices[i] ，s 的一个从 indices[i] 开始、
		到下一个 '#' 字符结束（但不包括 '#'）的 子字符串 恰好与 words[i] 相等
给定一个单词数组 words ，返回成功对 words 进行编码的最小助记字符串 s 的长度 。


前缀/后缀
如果单词 X 是 Y 的后缀，那么单词 X 就不需要考虑了，因为编码 Y 的时候就同时将 X 编码了。
例如，如果 words 中同时有 "me" 和 "time"，我们就可以在不改变答案的情况下不考虑 "me"。

如果单词 Y 不在任何别的单词 X 的后缀中出现，那么 Y 一定是编码字符串的一部分。

因此，目标就是保留所有不是其他单词后缀的单词，最后的结果就是这些单词长度加一的总和，
因为每个单词编码后后面还需要跟一个 # 符号。

算法
由数据范围可知一个单词最多含有 77 个后缀，所以我们可以枚举单词所有的后缀。
对于每个后缀，如果其存在 words 列表中，我们就将其从列表中删除。
为了高效删除，我们将 words 用哈希集合来存储。
class Solution {
    public int minimumLengthEncoding(String[] words) {
        Set<String> good = new HashSet<String>(Arrays.asList(words));
        for (String word: words) {
            for (int k = 1; k < word.length(); ++k) {
                good.remove(word.substring(k));
            }
        }

        int ans = 0;
        for (String word: good) {
            ans += word.length() + 1;
        }
        return ans;
    }
}
。
*/
func minimumLengthEncoding(words []string) int {
	out := 0
	m := map[string]bool{}
	for _, w := range words {
		m[w] = true  //存储word   标记 true
	}
	for w, _ := range m {  //key,val
		for i := 1; i < len(w); i++ {             //访问byte
			delete(m, w[i:])
		}
	}
	for w, _ := range m {
		out += len(w) + 1
	}
	return out
}

