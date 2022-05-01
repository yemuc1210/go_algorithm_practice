package lt1995

// import "sort"

/*
1995. 统计特殊四元组
给你一个 下标从 0 开始 的整数数组 nums ，返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：

nums[a] + nums[b] + nums[c] == nums[d] ，且
a < b < c < d

*/

//hash存储d-c  枚举a,b
func countQuadruplets(nums []int) (ans int) {
    cnt := map[int]int{}  //枚举
    for b := len(nums) - 3; b >= 1; b-- {  //末尾几个是可能的c d  逆枚举b
		//b 减少，c的取值范围逐渐增大 其余项不变
        for _, x := range nums[b+2:] {   //
			//将满足条件的b+1且d>c对应的nums[d]-nums[c]加入
            cnt[x-nums[b+1]]++
        }
		//枚举a
        for _, x := range nums[:b] {
            if sum := x + nums[b]; cnt[sum] > 0 {
                ans += cnt[sum]
            }
        }
    }
    return
}
 