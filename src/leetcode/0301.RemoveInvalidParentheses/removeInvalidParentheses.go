package lt301
/*   20.有效括号    22.括号生成
301. 删除无效的括号
给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。

返回所有可能的结果。答案可以按 任意顺序 返回。
*/
/*
注意到题目中要求最少删除，这样的描述正是广度优先搜索算法应用的场景，
并且题目也要求我们输出所有的结果。
我们在进行广度优先搜索时每一轮删除字符串中的 1 个括号，
直到出现合法匹配的字符串为止，此时进行轮转的次数即为最少需要删除括号的个数。

我们进行广度优先搜索时，每次保存上一轮搜索的结果，
然后对上一轮已经保存的结果中的每一个字符串尝试所有可能的删除一个括号的方法，
然后将保存的结果进行下一轮搜索。
在保存结果时，我们可以利用哈希表对上一轮生成的结果去重，从而提高效率。


*/
func isValid(str string) bool {
    cnt := 0
    for _, ch := range str {
        if ch == '(' {
            cnt++
        } else if ch == ')' {
            cnt--
            if cnt < 0 {
                return false
            }
        }
    }
    return cnt == 0
}

func removeInvalidParentheses(s string) (ans []string) {
    curSet := map[string]struct{}{s: {}}
    for {
        for str := range curSet {
            if isValid(str) {
                ans = append(ans, str)
            }
        }
        if len(ans) > 0 {
            return
        }
        nextSet := map[string]struct{}{}
        for str := range curSet {
            for i, ch := range str {
                if i > 0 && byte(ch) == str[i-1] {
                    continue
                }
                if ch == '(' || ch == ')' {
                    nextSet[str[:i]+str[i+1:]] = struct{}{}
                }
            }
        }
        curSet = nextSet
    }
}

