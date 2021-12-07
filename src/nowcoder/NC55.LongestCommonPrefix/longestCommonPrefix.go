package NC55

/**
  * 最长公共前缀 lt14
  * @param strs string字符串一维数组 
  * @return string字符串
*/
func longestCommonPrefix( strs []string ) string {
    // 1：排序？   2：做减法 从最长的串开始
	//二维数组  一列一列遍历  复杂度O(m*n)  费空间
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	//查找最长的串
	maxn,maxs := 0,""
	for _,str := range strs {
		if len(str) > maxn {
			maxn = len(str)
			maxs = str
		}
	}
	//从最长的字符串
	byteMaxStr := []byte(maxs)
	//做减法
	minCut := maxn    //maxn == len(byteMaxStr)
	//一个一个匹配
	for _,str := range strs {
		if str != maxs {
			//找到匹配的前缀
			idx := 0 

			for idx<len(str) {
				if str[idx] == byteMaxStr[idx] {  //匹配
					idx ++
				}else{
					if idx == 0 {
						return ""
					}
					break
				}
			}
			if idx < minCut {
				minCut = idx
			}
		}
	}
	return maxs[:minCut]
}