package day01

func twoSum(nums []int, target int) []int {
	//map
	m := make(map[int]int)
	for k,v := range nums {
		if idx,exist:=m[target - v]; exist {
			//存在配套的
			return []int{idx,k}
		}
		m[v] = k
	}
	return nil
}