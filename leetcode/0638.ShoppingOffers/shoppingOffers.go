package lt638



/*
638. 大礼包
在 LeetCode 商店中， 有 n 件在售的物品。每件物品都有对应的价格。然而，也有一些大礼包，每个大礼包以优惠的价格捆绑销售一组物品。

给你一个整数数组 price 表示物品价格，其中 price[i] 是第 i 件物品的价格。
另有一个整数数组 needs 表示购物清单，其中 needs[i] 是需要购买第 i 件物品的数量。

还有一个数组 special 表示大礼包，special[i] 的长度为 n + 1 ，
其中 special[i][j] 表示第 i 个大礼包中内含第 j 件物品的数量，
且 special[i][n] （也就是数组中的最后一个整数）为第 i 个大礼包的价格。

返回 确切 满足购物清单所需花费的最低价格，你可以充分利用大礼包的优惠活动。
你不能购买超出购物清单指定数量的物品，即使那样会降低整体价格。任意大礼包可无限次购买。

输入：price = [2,5], special = [[3,0,5],[1,2,10]], needs = [3,2]
输出：14
解释：有 A 和 B 两种物品，价格分别为 ¥2 和 ¥5 。
大礼包 1 ，你可以以 ¥5 的价格购买 3A 和 0B 。
大礼包 2 ，你可以以 ¥10 的价格购买 1A 和 2B 。
需要购买 3 个 A 和 2 个 B ， 所以付 ¥10 购买 1A 和 2B（大礼包 2），以及 ¥4 购买 2A 。


1.先计算只购买单品的金额，这部分用dfs的话太耗时。
2.开始对大礼包进行dfs+剪枝
- 剪枝1：如果礼包价格大于等于单品价格，放弃该礼包
- 剪枝2：选择礼包从前往后选，避免出现礼包1,2和礼包2,1的重复计算
- 剪枝3：如果礼包里的商品数量大于待购清单，放弃该礼包
3.如果选择完礼包后，还没有填满待购清单，后续通过购买单品完成，后续购买单品无需dfs操作，
直接按差量购买计算金额。


*/

func shoppingOffers(price []int, special [][]int, needs []int) int {
    cost := 0
    path := make([]int, len(price))
    //先计算购买单品的价格，如果这个部分用dfs的话，太耗时
    for k, v := range needs {
        cost += v * price[k]
    }
    dfs(price, special, needs, path, 0, &cost, 0)
    return cost
}
//path:当前购物总数；cur:当前消费金额; cost:最小消费金额；pos：当前礼包位置
func dfs(price []int, special [][]int, needs []int, path []int, cur int, cost *int, pos int) {
    if equal(needs, path) {
        if  cur < *cost {
            *cost = cur
        } 
        return
    }
    //choose 大礼包
    for k, v:= range special {
        if  cur+v[len(v)-1] >= *cost || k < pos { //不允许选择前面的礼包，避免出现礼包1 2, 2 1重复计算的情况
            continue
        }
        flag := true
        for k1, v1 := range v {
            if k1 < len(price) {
                if path[k1] + v1 > needs[k1] {
                    flag = false
                    break //如果礼包里的商品数量大于待购清单，放弃该礼包
                }
            }
        }
        if flag {
            for i:=0; i<len(price); i++ {
                path[i] = path[i] + v[i]
            }
            dfs(price, special, needs, path, cur+v[len(v)-1], cost, k)
            //cancle choose
            for i:=0; i<len(price); i++ {
                path[i] = path[i] - v[i]
            }
        }
    }
    //choose 单品
    for k, v  := range price {
        cur += (needs[k] - path[k]) * v
    }
    if cur < *cost {
        *cost = cur
    }
}
func equal(a []int, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for k, v := range a {
        if v != b[k] {
            return false
        }
    }
    return true
}
