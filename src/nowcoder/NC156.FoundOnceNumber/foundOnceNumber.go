package NC156

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * lt136  lt137
 * 数组中只出现一次的数 其他数出现k次 
 * @param arr int一维数组 
 * @param k int 
 * @return int
*/
func foundOnceNumber( arr []int ,  k int ) int {
    // write code here
	//模拟电路  lt137
	res := int32(0)
	for i:=0;i<32;i++{
		total := int32(0)
		for _,num := range arr{
			total += int32(num) >> i & 1
		}
		//答案的第 i 个二进制位就是数组中所有元素的第 i 个二进制位之和除以 3 的余数
		if total % int32(k) > 0 {
			res = res | 1<<i
		}
	}
	return int(res)
}