package lt917
/*
给你一个字符串 s ，根据下述规则反转字符串：

所有非英文字母保留在原有位置。
所有英文字母（小写或大写）位置反转。
返回反转后的 s 。
*/
func reverseOnlyLetters(s string) string {
    bs := []byte(s)
    for i,j := 0,len(bs)-1;i<j; {
        // 找到字符位置
        for i<len(bs) && !isLetter(bs[i])  {i++}
        for j>=0 && !isLetter(bs[j])  {j--}

        // 交换位置    7_28J
        if i<j {
            bs[i],bs[j] = bs[j],bs[i]
        }

        // 下一次迭代
        i++
        j--
    }
    return string(bs)
}

func isLetter(a byte) bool {
    return (a>='a'&&a<='z') ||(a>='A' && a<='Z')
}