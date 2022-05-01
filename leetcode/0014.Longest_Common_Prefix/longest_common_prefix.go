/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
package leetcode

// import "fmt"

func longestCommonPrefix(strs []string) string {
	//用树结构，首先找出最长的那个字符串构建树
	maxn,maxstr := 0,""
	for _,s := range strs{
		n := len(s)
		if n> maxn {
			maxn=n
			maxstr=s
		}
	}
	//根据maxstr 构建树    .....不用这么麻烦，数组也行啊
	bytes := []byte(maxstr)

	min_cnt := maxn
	for _,s := range strs{
		if s != maxstr {
			cnt_tmp := 0
			byte_tmps := []byte(s)

			for i,ch := range byte_tmps{
				// fmt.Printf("s=%s,ch=%v,bytes[i]=%v\n",s,byte_tmps[i:i+1],bytes[i:i+1])
				if ch==bytes[i]{
					cnt_tmp ++
				}else{
					if i==0{
						return ""
					}
					break
				}
				
			}
			// fmt.Printf("min_cnt=%v,cnt_tmp=%v\n",min_cnt,cnt_tmp)
			if cnt_tmp < min_cnt{
				min_cnt = cnt_tmp
			}
		}
	}
	// fmt.Printf("min_cnt=%v\n",min_cnt)

	return maxstr[:min_cnt]   //

}