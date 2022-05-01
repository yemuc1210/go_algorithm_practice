package lt443
/*
每日一题 
给你一个字符数组 chars ，请使用下述算法压缩：
从一个空字符串 s 开始。对于 chars 中的每组 连续重复字符 ：
如果这一组长度为 1 ，则将字符追加到 s 中。
否则，需要向 s 追加字符，后跟这一组的长度。
压缩后得到的字符串 s 不应该直接返回 ，需要转储到字符数组 chars 中。
需要注意的是，如果组长度为 10 或 10 以上，则在 chars 数组中会被拆分为多个字符。

请在 修改完输入数组后 ，返回该数组的新长度。
你必须设计并实现一个只使用常量额外空间的算法来解决此问题。


*/
func compress(chars []byte) int {
    //需要修改输入数组，再返回数组新长度
	//双指针   读指针  写指针
	write,left := 0 ,0
	for read,ch := range chars {
		if read == len(chars)-1 || ch != chars[read+1] {
			//read到了最右侧
			chars[write]=ch
			write++
			num := read - left + 1
			if num > 1{
				anchor := write
				for ;num >0; num/=10 {
					chars[write] = '0' + byte(num%10)   //写入数组，此时数组是倒着写入的
					write ++
				}
				s := chars[anchor:write]   //把数字倒置回来
				for i,n := 0,len(s); i<n/2; i++{
					//逆置
					s[i],s[n-i-1] = s[n-i-1],s[i]
				}
			}
			left = read+1 // 新字符的最左下标
		}
	}

	return write
    
}