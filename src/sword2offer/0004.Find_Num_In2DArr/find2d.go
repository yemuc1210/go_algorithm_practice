package offer04

func findNumberIn2DArray(matrix [][]int, target int) bool {
	//这种题有必要做二分法
		if len(matrix)==0 || len(matrix[0])==0{
			return false
		}
		//第一种，每行做二分，应该不是最优
		for i:=0;i<len(matrix);i++{
			if matrix[i][0] > target || matrix[i][len(matrix[i])-1] < target{
				continue
			}
			left,right := 0,len(matrix[i])-1
			for left <= right {
				mid := (left+right)/2
				// fmt.Println(mid,matrix[i][mid])
				if matrix[i][mid] == target{
					return true
				}else if matrix[i][mid] < target{
					left = mid + 1
				}else{
					right = mid -1
				}
			}
		}
	
		return false
	}