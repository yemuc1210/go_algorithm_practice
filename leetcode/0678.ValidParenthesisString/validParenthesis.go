package lt678

/*   lt20
678. 有效的括号字符串
给定一个只包含三种字符的字符串：（ ，） 和 *，写一个函数来检验这个字符串是否为有效字符串。
有效字符串具有如下规则：

任何左括号 ( 必须有相应的右括号 )。
任何右括号 ) 必须有相应的左括号 ( 。
左括号 ( 必须在对应的右括号之前 )。
* 可以被视为单个右括号 ) ，或单个左括号 ( ，或一个空字符串。
一个空字符串也被视为有效字符串。

*/
func checkValidString(s string) bool {
    var leftStk, asteriskStk []int
    for i, ch := range s {
        if ch == '(' {
            leftStk = append(leftStk, i)
        } else if ch == '*' {
            asteriskStk = append(asteriskStk, i)
        } else if len(leftStk) > 0 {   //否则，执行到这出现右括号，优先弹出左括号
            leftStk = leftStk[:len(leftStk)-1]
        } else if len(asteriskStk) > 0 {
            asteriskStk = asteriskStk[:len(asteriskStk)-1]
        } else {
            return false
        }
    }
    i := len(leftStk) - 1
    for j := len(asteriskStk) - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if leftStk[i] > asteriskStk[j] {
            return false
        }
    }
    return i < 0
}
