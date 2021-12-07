package lt229

/*
229. 求众数 II
给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

map ??
进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
*/

func majorityElement1(nums []int) []int{
	m := make(map[int]int)

	n := len(nums)
	res := []int{}
	for i:=0;i<n;i++{
		m[nums[i]]++
	}
	for k,v := range m{
		if v > n/3{
			res = append(res, k)
		}
	}
	return res
}

/*
异或  1^1=0   1^0=1 0^0=0
目标时间复杂度O(n)，空间复杂度O(1)
摩尔投票 同增异减
满足>n/3，这样元素最多两个
每次选择三个互不相同的元素删除
流程：
	每次检测当前元素是否是第一个或第二个选中的元素
	若都不相同，抵消一次
	最终存在的投票数大于0的，统计次数，检查
*/
func majorityElement(nums []int) (ans []int) {
    element1, element2 := 0, 0
    vote1, vote2 := 0, 0

    for _, num := range nums {
        if vote1 > 0 && num == element1 { // 如果该元素为第一个元素，则计数加1
            vote1++
        } else if vote2 > 0 && num == element2 { // 如果该元素为第二个元素，则计数加1
            vote2++
        } else if vote1 == 0 { // 选择第一个元素
            element1 = num
            vote1++
        } else if vote2 == 0 { // 选择第二个元素
            element2 = num
            vote2++
        } else { // 如果三个元素均不相同，则相互抵消1次
            vote1--
            vote2--
        }
    }

    cnt1, cnt2 := 0, 0
    for _, num := range nums {
        if vote1 > 0 && num == element1 {
            cnt1++
        }
        if vote2 > 0 && num == element2 {
            cnt2++
        }
    }

    // 检测元素出现的次数是否满足要求
    if vote1 > 0 && cnt1 > len(nums)/3 {
        ans = append(ans, element1)
    }
    if vote2 > 0 && cnt2 > len(nums)/3 {
        ans = append(ans, element2)
    }
    return
}

