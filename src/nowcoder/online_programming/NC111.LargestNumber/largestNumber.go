package NC111

import (
	"sort"
	"strconv"
)

/**
 * 最大数
 * @param nums int整型一维数组
 *@return string字符串
 */
func solve(nums []int) string {
	//根据最高位数字排序
	sort.Slice(nums, func(i, j int) bool {
		x,y := nums[i],nums[j]
		//获得最高位进行比较
		sx,sy := 10,10
		for sx <= x{
			sx *=10
		} //得到最高位的阶数
		for sy<=y {
			sy *=10
		}
		return sy*x+y > sx*y+x   //sy*x+y  x排前面   这里逆序（降序）
	})
	if nums[0] == 0 {
		//第一位最高的是0  其余也是0
		return "0"
	}
	res := []byte{}
	for _,num := range nums {
		res = append(res, strconv.Itoa(num)...)
	}
	return string(res)
} 