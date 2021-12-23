package NC142

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 定义重复字符串是由两个相同的字符串首尾拼接而成，
 * 例如 abcabcabcabc 便是长度为6的一个重复字符串，而 abcbaabcba 则不存在重复字符串。
 * 求最长重复子串的长度
 * @param a string字符串 待计算字符串
 * @return int整型
*/
func solve( a string ) int {
    slen := len(a) / 2
    for ; slen > 0; slen-- {
        Loop:for i := 0; i + 2 * slen <= len(a); i++ {
            for j := i;j < i + slen; j++ {
                if a[j] != a[j + slen]{
                    continue Loop;
                }
            }
            return slen * 2;
        }
    }
    
    return 0
}