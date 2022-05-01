package lt475

import (
	"math"
	"sort"
)

/*
475. 供暖器
冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。

在加热器的加热半径范围内的每个房屋都可以获得供暖。

现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，
请你找出并返回可以覆盖所有房屋的最小加热半径。

说明：所有供暖器都遵循你的半径标准，加热的半径也一样。
*/

/*
排序+二分
max heaters[i]<=houses

*/
func findRadius(houses, heaters []int) (ans int) {
    sort.Ints(heaters)
    for _, house := range houses {
        j := sort.SearchInts(heaters, house+1)
        minDis := math.MaxInt32
        if j < len(heaters) {
            minDis = heaters[j] - house
        }
        i := j - 1
        if i >= 0 && house-heaters[i] < minDis {
            minDis = house - heaters[i]
        }
        if minDis > ans {
            ans = minDis
        }
    }
    return
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/heaters/solution/gong-nuan-qi-by-leetcode-solution-rwui/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。