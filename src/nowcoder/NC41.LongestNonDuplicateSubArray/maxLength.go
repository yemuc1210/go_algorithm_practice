package NC41

// import "fmt"

/**
 * 最长无重复子数组  双指针+哈希
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxLength( arr []int ) int {
    // write code here
	if len(arr)==0 || len(arr)==1 {
		return len(arr)
	}
	res := 0
	left,right := 0,0  //左右指针
	
	newLoop: for left <= right && right < len(arr) {
		m := make(map[int]int)  //用map记录当前窗口出现次数
		m[arr[right]] ++  //起点
		// right ++
		for right=left+1; right < len(arr); right++{
			// fmt.Println(m,right, arr[right],m[arr[right]])
			if _,ok := m[arr[right]]; !ok {
				//还未存入
				m[arr[right]] += 1
			}else{ //否则存在重复元素
				res = max(res, right-left)
				left = right
				goto newLoop
			}
		}
	}
	//可能整个数组都是无重复的  所以res没有在循环内部更新
	res = max(res, right-left)
	return res
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}