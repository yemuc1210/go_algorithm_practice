package NC118

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 数组中的逆序对   offer51
 * @param data int整型一维数组 
 * @return int整型
*/
func InversePairs( data []int ) int {
    // 一般思路：每一个数字都往后遍历一次   复杂度太高
	// 要求：空间复杂度O(n)  时间复杂度O(nlogn)  
	//排序 得到交换次数 就是逆序对个数  排序应该选
	//P mod 1000000007
	return mergeSort(data, 0, len(data)-1) % 1000000007
}

func mergeSort(nums []int, start,end int) int {
	if start >= end {
		return 0
	}  //区间是[start,end]
	mid := start + (end-start)>>2
	left := mergeSort(nums, start, mid)   //左边的交换次数  [start,mid]
	right := mergeSort(nums, mid+1, end)   //右边的交换次数 [mid+1, right]
	cross := merge(nums,start,mid,end)
	return left+right+cross
}

func merge(nums []int, start,mid,end int) int {
	arr := make([]int, end-start+1)
	p, q, k, count := start, mid+1, 0, 0
	for i:=start;i<=end;i++{
		if p>mid {
			arr[k] = nums[q]
			q++
		}else if q>end{
			arr[k] = nums[p]
			p++
		}else if nums[p] <= nums[q] {
			arr[k] = nums[p]
			p++
		}else{  //nums[p] > nums[q]
			//发生交换 
			count += mid - p + 1   //因为数组[p,mid]都是有序的 且递增的
			arr[k] = nums[q]
			q++
		}
		k++
	}
	//将归并结果存原来的切片
	copy(nums[start:end+1], arr)
	return count
}

