package lt414
/*
简单题

414. 第三大的数
给你一个非空数组，返回此数组中 第三大的数 。
如果不存在，则返回数组中最大的数。

要求O(n)


示例 3：

输入：[2, 2, 3, 1]
输出：1
解释：注意，要求返回第三大的数，是指在所有不同数字中排第三大的数。
此例中存在两个值为 2 的数，它们都排第二。在所有不同数字中排第三大的数为 1 。

*/

/*
我们可以遍历数组，同时用一个有序集合来维护数组中前三大的数。具体做法是每遍历一个数，就将其插入有序集合，若有序集合的大小超过 33，就删除集合中的最小元素。这样可以保证有序集合的大小至多为 33，且遍历结束后，若有序集合的大小为 33，其最小值就是数组中第三大的数；若有序集合的大小不足 33，那么就返回有序集合中的最大值。
*/
func thirdMax(nums []int) int {
    t := redblacktree.NewWithIntComparator()
    for _, num := range nums {
        t.Put(num, nil)
        if t.Size() > 3 {
            t.Remove(t.Left().Key)
        }
    }
    if t.Size() == 3 {
        return t.Left().Key.(int)
    }
    return t.Right().Key.(int)
}


func thirdMax(nums []int) int {
    a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
    for _, num := range nums {
        if num > a {
            a, b, c = num, a, b
        } else if a > num && num > b {
            b, c = num, b
        } else if b > num && num > c {
            c = num
        }
    }
    if c == math.MinInt64 {
        return a
    }
    return c
}



