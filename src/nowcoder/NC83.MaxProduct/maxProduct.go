package NC83

/**
 * 连续子数组最大乘积
 * @param arr double浮点型一维数组 
 * @return double浮点型
*/
func maxProduct( arr []float64 ) float64 {
    // write code here
	n := len(arr)
	if n==0 {
		return 0
	}
	if n==1{
		return arr[0]
	}
	minimum,maximum,res := arr[0],arr[0],arr[0]
	for i:=1;i<n;i++{
		if arr[i] <0 {
			//对结果不利
			maximum,minimum = minimum,maximum
		}
		maximum = max(arr[i], arr[i]*maximum)
		minimum = min(arr[i], arr[i]*minimum)
		res = max(res,maximum)
	}
	return res
}

func max(a,b float64) float64{
	if a>b {
		return a
	}
	return b
}

func min(a,b float64) float64{
	if a<b {
		return a
	}
	return b
}