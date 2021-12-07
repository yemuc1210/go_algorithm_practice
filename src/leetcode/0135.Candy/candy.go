package lt135

// import "fmt"

/*困难
135. 分发糖果
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
你需要按照以下要求，帮助老师给这些孩子分发糖果：
	每个孩子至少分配到 1 个糖果。
	评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？
*/
/*
思路1：首先找到“谷底” 给其分配1个糖果
	特殊示例 【1，2，2】  分配方案是【1，2，1】
	（1）.找到谷底 分配1
	（2）根据规则：
		评分连续降低的，除了第一个是2 其余都是1；
		连续相等的 只有第一个是2 其余全是1
		评分连续上升的，除了最后一个是1 ，其余全是2
思路2：贪心  求最优值的嘛
	先确定一边，再确定另一边  局部最优
	先确定右边评分大于左边的   【从前向后遍历】
		此时是局部最优，只要右边比左边大，右边就多一个
		全局最优：相邻孩子中，评分高的右比左多
	接着确定左大于右的情况  【从后往前遍历】
		rank[i] > rank[i+1] 则i有两个选择：
			candy[i+1]+1   ||   cnady[i]
		贪心局部最优 取较大值
*/
func candy(ratings []int) int {
	cadnys := make([]int, len(ratings))
	// fmt.Println(cadnys)
	//初始化1
	for i:=range cadnys {
		cadnys[i] = 1
	}
	//从前往后遍历
	for i:=1;i<len(ratings);i++{
		if ratings[i] > ratings[i-1] {
			cadnys[i] = cadnys[i-1] + 1
		}
	}
	//end - front
	for i:=len(ratings)-2;i>=0;i-- {
		if ratings[i] > ratings[i+1] {
			cadnys[i] = max(cadnys[i], cadnys[i+1]+1)
		}
	}
	//get result
	res := 0 
	for i:=0;i<len(ratings);i++{
		res += cadnys[i]
	}
	return res
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}