package NC74

/**
 * 升序数组中数字k出现次数
 * @param data int整型一维数组 
 * @param k int整型 
 * @return int整型
*/
func GetNumberOfK( data []int ,  k int ) int {
    if len(data) == 0 {
		return 0
	}
	var cnt int
	for i:=0;i<len(data);i++{
		if data[i] == k {
			//从i开始若干个都是K
			cnt =1
			for j:=i+1;j<len(data);j++{
				if data[j] == k {
					cnt ++
				}else{
					return cnt
				}
			}
			return cnt
		}
	}
	return cnt
}