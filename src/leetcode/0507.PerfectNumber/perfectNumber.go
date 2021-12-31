package lt507

func perfectNumer(num int) bool {
	//完美数：除了它本身 = 自身以外的所有正因子之和相等
	//如 28 = 1+2+4+7+14
	if num <= 1 {
		return false
	}
	sum := 1 //肯定包含因子1
	for i:=2; i*i<=num;i++{
		if num%i==0 {
			//i是一个因子
			sum += i
			if i*i<num{
				sum += num/i
			}
		}
	}
	return sum==num
}