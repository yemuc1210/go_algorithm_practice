package leetcode

import "sort"

/*
夏日炎炎，小男孩 Tony 想买一些雪糕消消暑。
商店中新到 n 支雪糕，用长度为 n 的数组 costs 表示雪糕的定价，
其中 costs[i] 表示第 i 支雪糕的现金价格。Tony 一共有 coins 现金可以用于消费，
他想要买尽可能多的雪糕。
给你价格数组 costs 和现金量 coins ，
请你计算并返回 Tony 用 coins 现金能够买到的雪糕的 最大数量 。
注意：Tony 可以按任意顺序购买雪糕。
*/

// type pair struct{
// 	curCost int
// 	curLeft int
// 	coinsLeft int
// }
// func maxIceCream(costs []int, coins int) int {
// 	sum:=0
// 	for _,v:=range costs{
// 		sum += v
// 	}
// 	if sum<=coins{
// 		return len(costs)
// 	}
// 	left := sum     //剩余
// 	min := pair{curCost: costs[0],curLeft: sum-costs[0],coinsLeft: coins-costs[0]}

// 	//尝试使用剪枝的方法写
// 	queue := []pair{min}
// 	for len(queue)!=0 {
// 		cur := queue[0]
// 		queue = queue[1:] //出队
// 		if cur.curCost < min.curCost {

// 		}
// 	}

// }
//排序+贪心
func maxIceCream(costs []int, coins int) (ans int) {
    sort.Ints(costs)   //O(nlogn)
    for _, c := range costs {
        if coins < c {
            break
        }
        coins -= c
        ans++
    }
    return
}
//方法二  剪枝回溯
/*
剪枝条件？？好像不行，剪枝的前提也是排序
如果非要排序，那么直接配合贪心好了

空间换时间也行，新建数组
｛price,nums,sumprice,left｝
记录总价，以及选择购买之后可以剩余的coins，贪心选择剩余最大的
感觉还是会排序
*/
// func maxIceCream1(costs []int,coins int)int{
// 	return 0
// }