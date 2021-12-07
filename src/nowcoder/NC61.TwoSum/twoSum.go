package NC61

/**
  * 两数之和
  * @param numbers int整型一维数组 
  * @param target int整型 
  * @return int整型一维数组
*/
func twoSum( nums []int ,  target int ) []int {
    // 找出一对
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx+1, k+1}
		}
		m[v] = k
	}
	return nil
}

func factor(n int) int{
	if n==1 || n==0{
		return 1
	}
	return n*factor(n-1)
}