package main

import (
	"fmt"
	"math"
)

// func nextBeautifulNumber(n int) int {
// 	// x: 每个数位d，恰好在x中出现d次，则x是数值平衡数
// 	// 给定n，返回严格大于n的的最小数值平衡数
// 	// 打表
// 	// 1 22 333 4444 55555 666666 7777777   这几个数组合
// 	// 组合：122 1333  221  3331  得到所有组合结果，然后sort
// 	origins := []string{"1", "22", "333", "4444", "55555", "666666", "7777777"}
// 	// 组合  求组合使用dfs
// 	sets := []string{}
// 	dfs(0, len(origins), &sets, []string{}, origins)
// 	fmt.Println(len(sets))
// 	m   := make(map[int]int)
// 	nums := []int{}
// 	for _,v := range sets {
// 		num,_ := strconv.Atoi(v)
// 		m[num]++
// 	}
// 	for _,v := range origins {
// 		num,_ := strconv.Atoi(v)
// 		m[num]++
// 	}
// 	for k := range m {
// 		nums = append(nums, k)
// 	}
// 	sort.Ints(nums)
// 	fmt.Println(nums)
// 	// 二分查找
// 	res := search(nums,n)
// 	return nums[res+1]
// }

// func search(nums []int, target int) int {
// 	left,right := 0,len(nums)-1
// 	for left<=right {
// 		mid := left +(right-left)/2
// 		if nums[mid] == target {
// 			return mid
// 		}else if nums[mid] > target {
// 			right = mid-1
// 		}else{
// 			left = mid+1
// 		}
// 	}
// 	return right
// }
// func dfs(curIdx int, n int, res *[]string, path []string, origins []string) {
// 	if curIdx == n {
// 		// 得到一组结果
// 		tmp := make([]string, len(path))
// 		copy(tmp, path)
// 		*res = append(*res, tmp...)
// 		return
// 	}

// 	//否则，origins[curIdx] 与其他数字匹配
// 	for i:=0;i<n;i++ {
// 		if i==curIdx{
// 			continue
// 		}
// 		path = append(path, origins[curIdx]+origins[i])
// 		dfs(curIdx+1, n, res, path, origins)
// 		path = path[:len(path)-1]
// 	}
// }
// 线性遍历
// 数据范围限定
func nextBeautifulNumber(n int) int {
	// 从n+1开始遍历
	if n==0 {return 1}
	i:=n+1
	for ;i<math.MaxInt32;i++{
		if isBalanced(i) {
			break
		}
	}
	return i
}

func isBalanced(n int) bool{
	// 哈希统计1-7的出现次数
	hash := make([]int,8)

	// 统计
	for n>0 {
		if n%10>7 || n%10==0 {
			return false
		}
		hash[n%10] ++
		n/=10
	}

	// 只要出现次数不为0且不等于自身，则不是平衡
	for i:=1;i<=7;i++{
		if hash[i]!=0 && hash[i]!=i {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(nextBeautifulNumber(3000))
}
