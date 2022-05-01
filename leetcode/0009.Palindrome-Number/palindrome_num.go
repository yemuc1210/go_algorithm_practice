package leetcode

import (
	// "fmt"
	"strconv"
)

// import "fmt"

/*
判断数字是否是回文
输入：121
输出：true

-121
flase     121-
*/
func palidreome_number(number int)bool{
	
	if number < 0 {
		return false
	}
	if number == 0 {
		return true
	}
	if number%10 == 0 {
		return false
	}
	
	// var res bool
	//那么现在就处理整数，起码两位数
	
	a := number
	new_num := 0
	for a>0 {
		new_num = new_num * 10 +a % 10
		a = a/10
	}


	// fmt.Printf("new_num=%v,number=%v",new_num,number)
	
	return new_num == number
}

//解法二：转化成字符串
func palidreome_number_by_string(x int) bool{
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}