package lt859

import "reflect"

/*简单 每日一题
859. 亲密字符串
给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，就返回 true ；否则返回 false 。

交换字母的定义是：取两个下标 i 和 j （下标从 0 开始）且满足 i != j ，接着交换 s[i] 和 s[j] 处的字符。

例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。
提示：
	1 <= s.length, goal.length <= 2 * 104
	s 和 goal 由小写英文字母组成
*/
/*
思路1：
	一定false的情况：长度不想等、字符出现频次不相同
		统计字符出现频次
	做字符串减法 出现两个非零且两个数和为0 才行
		因为根据提示  只有小写字母出现
	错误：无法识别两个可交换数字是不是aa这类情况

	完全相等：查找有无重复的情况 有则true  无则false
	不相等：需要记录差值及对应字符  定义结构体

*/
func buddyStrings(s string, goal string) bool {
	if len(s)-len(goal) != 0 {
		return false
	}
	diff := make([]int,len(s))  //字符减法
	//做减法 
	for i:=0;i<len(s);i++{
		diff[i] = int(s[i]-'a') - int(goal[i]-'a')
	}
	//遍历diff  最多只能出现两个非零数 且 和为0
	//不排除完全相等的情况  根据示例  ab  ab 为 false   aa  aa 为true
	type pair struct{
		ch1,ch2 byte
		diff int
	}
	nonZeor := []pair{}
	for i:=range diff{
		if diff[i] != 0 {
			nonZeor = append(nonZeor, pair{s[i],goal[i],diff[i]})
		}
	}
	if len(nonZeor)==0{
		//查找有无重复 即aa这种
		m := make(map[rune]int)
		for _,ch := range s{
			m[ch]++
		}
		//遍历m  有重复 则某个字符出现偶数次
		for _,v := range m{
			if v>=2 {
				return true
			}
		}
		return false
	} 
	if len(nonZeor)==2 {
		// 单纯比较差值不行   如b-c  b-a = -1 1  记录这两队字符
		//和0  字符是否匹配
		if nonZeor[0].ch1==nonZeor[1].ch2 && nonZeor[0].ch2==nonZeor[1].ch1 && nonZeor[0].diff+nonZeor[1].diff==0 {
			return true
		}
		return false
	}
	return false
}

//范例
func buddyStrings1(s string, goal string) bool {
    if len(s) != len(goal){
        return false
    }
    n := len(s)
    a := []byte{}
    b := []byte{}
    for i := 0; i < n; i++{
        if s[i] != goal[i]{
            a = append(a, s[i])
            b = append(b, goal[i])
        }
    }

    if len(a) == 2 && len(b) == 2 && a[0] == b[1] && a[1] == b[0] {
        return true
    }
    if reflect.DeepEqual(s, goal) {
        alphaMap := map[rune]int{}
        for _, s := range s {
            alphaMap[s]++
            if alphaMap[s] == 2{
                return true
            }
        }
    }

    return false
}   