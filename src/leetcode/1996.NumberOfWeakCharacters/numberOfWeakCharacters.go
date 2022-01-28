package lt1996

import "sort"

func numberOfWeakCharacters(properties [][]int) int {
	//pro[i] = [attack_i, defenses_i]
	// 根据一个指标降序  根据攻击力
	sort.Slice(properties, func(i, j int) bool {
        // 攻击力降序  若相同则防御力升序
        if properties[i][0] == properties[j][0] {
            return properties[i][1] < properties[j][1]
        }
		return properties[i][0] > properties[j][0]
	})

	//遍历 
	res := 0
	curMax := -1   // 已经遍历过的防御力的最大值
	for i:=0;i<len(properties);i++{
		if curMax > properties[i][1] {
			res ++  // 攻击防御皆比其高
		}
		curMax = max(curMax, properties[i][1])
	}
	return res
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}
//手写一个排序