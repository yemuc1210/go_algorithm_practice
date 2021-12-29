package NC89

/**
 * go + 两次翻转
先转换成byte切片
大小写转换
整体翻转
找出每一个单词进行翻转
 * @param s string字符串 
 * @param n int整型 
 * @return string字符串
*/
func trans( s string ,  n int ) string {
    // write code here
    if n == 0 {
        return s
    }

    bs := []byte(s)
//  对每一个byte进行大小写转换
    for i, x := range bs {
        bs[i] = transByte(x)
    }

//  先整体翻转
    for i, j := 0, n-1; i < j; i, j = i+1, j-1{
        bs[i], bs[j] = bs[j], bs[i]
    }


//  单词的开始、结束位置
    start, end := 0, 0

//  再次翻转每一个单词
    for  end < n {
        for end < n && bs[end] == ' ' {
            end++
        }
        start = end

//      找到一个单词
        for end < n && bs[end] != ' ' {
            end++
        }

//      翻转单词
        reverseStr(bs, start, end-1)

//      找到空格
        for end < n && bs[end] == ' ' {
            end++
        }

//      到这里已经是新的单词位置了
        start = end
    }

//  最后一个单词 这里 end == n, 所以需要end-1
//  单词翻转  start <= end-1 才说明需要翻转
    if start <= end-1 {
        reverseStr(bs, start, end-1)
    }

    return string(bs)
}

// 位置交换
func reverseStr(bs []byte, s, e int) {
    for ; s < e ; s, e = s+1, e-1 {
        bs[s], bs[e] = bs[e], bs[s]
    }
}

// 大小写转换
func transByte(b byte) byte{
    if b >= 'a' && b <= 'z' {
        return  b - ('a'-'A')
    }else if b >= 'A' && b <= 'Z'{
        return b + ('a'-'A')
    }
    return b
}