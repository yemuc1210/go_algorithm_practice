package interview16


func calculator(s string ) int{
	// 符号包括：正整数  加减乘除 
	// 计算器中缀转后缀
	nums := []int{}
	curNum:=0
	preOp :='+'
	s = s+"+"
	for i:=0;i<len(s);i++{
		if s[i]>='0' && s[i]<='9' {
			curNum = curNum*10 + int(s[i]-'0')
		}else if s[i]==' ' {
			continue
		}else {
			if preOp == '+' {
				nums = append(nums, curNum)
			}else if preOp =='-' {
				nums = append(nums, -curNum)
			}else if preOp == '*' {
				nums[len(nums)-1] = nums[len(nums)-1]*curNum
			}else if preOp == '/'{
				nums[len(nums)-1] = nums[len(nums)-1]/curNum
			}
			preOp = rune(s[i])
			curNum = 0
		}
	}
	//栈做加法
	sum := 0 
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	return sum
}

