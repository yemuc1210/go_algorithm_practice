package offerS19
/*
剑指 Offer II 019. 最多删除一个字符得到回文
给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。

 
*/
func validPalindrome(s string) bool {
    left,right := 0,len(s)-1

    for left < right {
        //判断两个字符是否相等
        if s[left] == s[right] {
            left++
            right--
        }else{//否则需要考虑删除一个字符，并且只能删除一个
            leftIsPalin,rightIsPalin := true,true
            //删除左边的字符   s[left+1 : right]
            for i,j := left+1,right; i<j ; i,j=i+1,j-1{
                if s[i]!=s[j]{
                    leftIsPalin=false
                    break
                }
            }        

            //删除右边的字符   s[left : right-1]
            for i,j := left,right-1; i<j; i,j=i+1,j-1 {
                if s[i] != s[j] {
                    rightIsPalin = false
                    break
                }
            }
            return leftIsPalin || rightIsPalin
        }
    }
    return true
}