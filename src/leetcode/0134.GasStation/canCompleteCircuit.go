package lt134
/*
134. 加油站
在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。
你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
说明: 
	如果题目有解，该答案即为唯一答案。
	输入数组均为非空数组，且长度相同。
	输入数组中的元素均为非负数。
示例 1:

输入: 
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
输出: 3
解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。

输入: 
gas  = [2,3,4]
cost = [3,4,3]
输出: -1
*/
/*
gas[i]-cost[i]   即station i的有效加油
	gas  = [1,2,3,4,5]
	cost = [3,4,5,1,2]
	effective = [-2,-2,-2,3,3]   

	gas  = [2,3,4]
	cost = [3,4,3]
	effective=[-1,-1,1]   false
*/
/*
首先计算有效加油数组，根据有效加油数组得到备选解
	可能会超时，所以可以需要对备选解进行预判
	
然后遍历备选解得到唯一解。需要计算
*/
func canCompleteCircuit(gas []int, cost []int) int {
	effective := []int{}  //记录有效加油站的下标  即备选解
	sumEffective := 0
	for i:=0;i<len(gas);i++{
		if gas[i] - cost[i] >= 0 {
			effective = append(effective, i)
		}
		sumEffective += gas[i] - cost[i]
	}
	if sumEffective < 0 {
		return -1
	}
	res := -1
	//遍历备选解，得到唯一解 题目提示了若有解则是唯一解
	for i:=0;i<len(effective);i++{
		if isValidAnswer(gas,cost, effective[i]) {
			res = effective[i]
			break
		}
	}
	return res
}

func isValidAnswer(gas []int, cost[] int, startIdx int) bool {
	//从当前下标开始
	n := len(gas)
	idx := (startIdx+1)%n  //从startIdx开始，一定能到下一站
	curLeftGas := gas[startIdx] - cost[startIdx]  //到下一站剩余油量
	for idx != startIdx {  //判断能否回到startIdx
		curLeftGas = curLeftGas + gas[idx] - cost[idx]  
		if curLeftGas >= 0 {
			idx = (idx + 1) % n
		}else{
			break
		}
		
	}
	return idx == startIdx
}	

/*
【笔记】一次遍历法，车能开完全程需要满足两个条件：
	车从i站能开到i+1。
	所有站里的油总量要>=车子的总耗油量。
那么，假设从编号为0站开始，一直到k站都正常，在开往k+1站时车子没油了。
	这时，应该将起点设置为k+1站。
*/
func canCompleteCircuit1(gas []int, cost []int) int {
	rest,run,start := 0,0,0

	for i:=0;i<len(gas);i++{
		run += (gas[i] - cost[i])
		rest += (gas[i] - cost[i])
		if run < 0 {
			start = i+1
			run = 0
		}
	}
	if rest < 0 {  //即有效effective的概念 如果整体的有效值为负 一定不行
		return -1
	}
	return start
}