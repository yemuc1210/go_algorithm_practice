package NC130

//lt135
//分糖果：每个孩子最少一个 得分高的比两侧都多  最少需要多少个糖果
func candy(ratings []int) int {
	cadnys := make([]int, len(ratings))
	// fmt.Println(cadnys)
	//初始化1
	for i:=range cadnys {
		cadnys[i] = 1
	}
	//从前往后遍历
	for i:=1;i<len(ratings);i++{
		if ratings[i] > ratings[i-1] {
			cadnys[i] = cadnys[i-1] + 1
		}
	}
	//end - front
	for i:=len(ratings)-2;i>=0;i-- {
		if ratings[i] > ratings[i+1] {
			cadnys[i] = max(cadnys[i], cadnys[i+1]+1)
		}
	}
	//get result
	res := 0 
	for i:=0;i<len(ratings);i++{
		res += cadnys[i]
	}
	return res
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}