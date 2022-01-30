package NC105

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 二分查找 含有重复元素的升序数组 lt34
 * 如果目标值存在返回下标，否则返回 -1
 * @param nums int整型一维数组 
 * @param target int整型 
 * @return int整型
*/
func search( nums []int ,  target int ) int {
    // write code here
    if len(nums)==0{
        return -1
    }
	start,end := 0,0
	if nums[0]==target{
		return start
	}
	for i:=end;i<len(nums);{
		if nums[i]==target{
			start = i		
			j:=i
			for ;j<len(nums)&&nums[j]==target;j++{}
			end = j-1
			i = j
			return start
		}else{
			i++
		}
	}

	return -1
}