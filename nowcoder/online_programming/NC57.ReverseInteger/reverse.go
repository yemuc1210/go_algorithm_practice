package NC57

/**
  * 反转数字  有符号整数 只反转数字部分
  * 反转后数字不能超出32位有符号整数范围
  * @param x int整型 
  * @return int整型
*/
func reverse( x int ) int {
    // write code here
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}