package offer15

func hammingWeight(num uint32) int {
    cnt := 0
	one := uint32(1)
	for num != 0{
		if num & one != 0{   //最后一位是1   
			cnt ++
		}
		num = num >> 1
	}
	
	return cnt
}