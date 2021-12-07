package lt881

import "sort"

/*救生艇
第 i 个人的体重为 people[i]，每艘船可以承载的最大重量为 limit。

每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。

返回载到每一个人所需的最小船数。(保证每个人都能被船载)。

示例 1：

输入：people = [1,2], limit = 3
输出：1
解释：1 艘船载 (1, 2)

*/
func numRescueBoats(people []int, limit int) int {
	//贪心法
	//首先排序
	sort.Ints(people)
	//双指针
	front,rear := 0,len(people)-1

	nums := 0

	for front <= rear{
		if people[rear] > limit{
			//那么只能单独载当前最重的
			rear --
			nums ++
		}else {
			//此时
			if people[front] + people[rear] <= limit {
				front ++
				rear --
				nums ++
			}else {
				//否则，仍然只能单独载最重的
				rear --
				nums ++
			}
		}
	}
	return nums
}