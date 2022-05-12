package main

// 字符串都是小写，且长度都想等
func minDeletionSize(strs []string) int {
	size,length := len(strs),len(strs[0])
	// 不是按字典升序排列的  列，删除
	// 所有构造二维数组，判断
	arr := [][]byte{} 
	for _,str:=range strs {
		arr = append(arr, []byte(str))
	}
	// 判断
	var res int 
	for i:=0;i<length;i++ {
		var num byte
		for j:=0;j<size;j++{
			if j==0 {
				num = arr[j][i]
				continue
			}
			if num <= arr[j][i] {
				num = arr[j][i]
				continue
			}else{
				// 需要删除
				res ++
				break  // 跳到下一列
			}
		}
	}
	return res
}