package offer62
//约瑟夫环问题
func lastRemaining(n int, m int) int {
    idx := 0
    for i:=2; i<=n; i++ {
        idx = (idx+m)%i
    }
    return idx
}