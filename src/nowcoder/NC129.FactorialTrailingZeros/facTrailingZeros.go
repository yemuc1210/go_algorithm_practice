package NC129
/**
 * the number of 0
 * 阶乘末尾零的个数
 * 有5的阶乘才有0
 * @param n long长整型 the number
 * @return long长整型
*/
func thenumberof0( n int64 ) int64 {
    // write code here
	if n/5 == 0 {
		return 0
	}
	return n/5 + thenumberof0(n/5)
}