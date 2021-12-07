package leetcode

// import "fmt"
/*
已知罗马字母表示数字

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

现在输入字母组合，求对应数字
如II=2
XII=10+2=12
XXVII  27
III  3
IV   4
VI  6
IX  9
XI  11
CD  400

左大右小  加法
左小右大   减法
*/

func  roman2int(roman string)int{
	m := map[string]int {
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	}   // 映射关系

	if roman == ""{
		return 0
	}
	start,end := 0 ,len(roman)-1
	// factor := 1  //乘子，控制符号
	res := 0
	for start < end {
		ch := string(roman[start])
		ch1 := string(roman[start+1])
		if m[ch1]<=m[ch] {
			//说明当前符号满足左大右小
			
			res = res + m[ch]
			// fmt.Println(res,m[ch])
			start++
		}else{ //否则这两个元素应该为一组   是那种
			// factor *= -1
			res = res + m[ch] * -1 + m[ch1]
			start += 2
		}
		
	}
	// fmt.Printf("out start=%v,end=%v,res=%v\n",start,end,res)
	if start == end{
		res += m[string(roman[end])]	
	}
	return res
	/*
	执行用时：8 ms, 在所有 Go 提交中击败了73.03%的用户
	内存消耗：3 MB, 在所有 Go 提交中击败了80.02%的用户
	*/
}

// func roman2int_1(s string)int{
// 	//换个思路，从右往左，如果是两个为一组需要特别关注的
// }
func romanToInt(s string) int {
	var roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastint {
			total = total - num
		} else {
			total = total + num
		}
		lastint = num
	}
	return total
}