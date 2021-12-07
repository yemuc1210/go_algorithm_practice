package interview01



/*
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

不适用额外的数据结构

假设字符范围是26个字母，否则是unicode的 2^16/8字符数


方法一：排序
方法二：map计数

*/
func isUnique(astr string) bool {
	//使用map
	m := make(map[rune]int, len(astr))

	for _,v := range astr{
		m[v] += 1
		if m[v] > 1{
			return false
		}
	}
	return true
}

//最好的应该还是位运算
/*
把不同的字符与字符'a'之间的差值，转换成1左移的位数，
比如字符'c','c'-'a'=2,把1左移两位 ==> 100，d就是1000
把差值与一个数num进行与运算(&),num初始值为0，

如果之前没有相同的字符出现，&结果为0，把出现的字符用或运算(|)存储到num中，以便下一次比较
如果之前有相同的字符出现，&结果为1

*/
func isUnique1(astr string) bool {
	num := 0
	for _,v := range astr{
		moveBit := v - 'a'
		if num & (1 << moveBit)!=0 {  //即有相同元素
			return false
		}else{
			num = num | (1<<moveBit)
		}
	}
	return true

}
