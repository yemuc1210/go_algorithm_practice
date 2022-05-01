package offer51



/*

在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
输入一个数组，求出这个数组中的逆序对的总数。

示例 1:
输入: [7,5,6,4]
输出: 5

注意输出总数
做一个排序，得到交换的次数
*/
// func reversePairs(nums []int) int {
// 	if len(nums)==0 || len(nums)==1{
// 		return len(nums)
// 	}
// 	n := len(nums)
// 	tmpArr := make([]int,n)
// 	res := 0
// 	mergeSort(nums,&tmpArr,0,n-1,&res)
// 	fmt.Println(tmpArr)
// 	return res
// }

// func merge(a []int,tmpArr *[]int,start,end,mid int,cnt *int){
// 	// tmpArr := make([]int,len(a)+len(b))
// 	i,j,k := start,mid+1,start
// 	for i<mid+1 && j<end+1{
// 		if a[i] > a[j]{  //从小到大排列
// 			(*tmpArr)[k] = a[j]
// 			k++
// 			j++
// 			*cnt = *cnt + mid - i + 1
// 			fmt.Printf("cnt=%v",*cnt)
// 		}else{
// 			(*tmpArr)[k] = a[i]
// 			k++
// 			i++
// 		}
// 	}
// 	for i<=mid{
// 		(*tmpArr)[k]=a[i]
// 		k++
// 		i++
// 	}
// 	for j<=end{
// 		(*tmpArr)[k]=a[j]
// 		k++
// 		j++
// 	}
	
// }
// func mergeSort(a []int,res *[]int,start,end int,cnt *int){
// 	if start<end{
// 		mid := (start+end)/2
// 		mergeSort(a,res,start,mid,cnt)
// 		mergeSort(a,res,mid+1,end,cnt)
// 		merge(a,res,start,mid,end,cnt)
// 		fmt.Println(res)
// 	}
// }

func reversePairs(nums []int) int {
    return merge_sort(nums, 0, len(nums)-1)
}

func merge_sort(A []int, start, end int) int {
    if start >= end {
        return 0
    }
    mid := start + (end - start)>>1
    left := merge_sort(A, start, mid)
    right := merge_sort(A, mid+1, end)
    cross := merge(A, start, mid, end)
    return left + right + cross
}
func merge (A []int, start, mid, end int) int{
    Arr := make([]int, end-start+1)
    p, q, k, count := start, mid+1, 0, 0
    for i := start; i <= end; i++ {
        if p > mid {
            Arr[k] = A[q]
            q++
        } else if q > end {
            Arr[k] = A[p]
            p++
        } else if A[p] <= A[q] {
            Arr[k] = A[p]
            p++
        } else {
            count += mid - p + 1
            Arr[k] = A[q]
            q++
         }
        k++
    }
    copy(A[start:end+1], Arr)
    return count
}
