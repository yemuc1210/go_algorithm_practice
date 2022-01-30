package NC56

/**
  * 回文数字
  * @param x int整型 
  * @return bool布尔型
*/
func isPalindrome( number int ) bool {
    // 不使用额外内存空间的条件下判断整数是否是回文   
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