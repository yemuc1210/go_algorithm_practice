package offer11

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。

例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  

*/

//递增排序的数组  经过旋转
func minArray(numbers []int) int {
	if len(numbers)==0{
		return -1
	}
	if len(numbers)==1{
		return numbers[0]
	}
	//现在，数组起码有两个元素
	for i:=0;i<len(numbers);i++{
		if i+1<len(numbers) && numbers[i+1] < numbers[i]{
			return numbers[i+1]
		}
	}
	return numbers[0]
}

/*既然是查找，试试二分查找
数组中可能存在重复的元素
旋转之后的数组，较大的元素在前面，较小的元素在后面
*/
func minArray1(numbers []int)int{
	if len(numbers)==0{
		return -1
	}
	if len(numbers)==1{
		return numbers[0]
	}
	low, high := 0, len(numbers)-1
	for low < high {
		if numbers[low] < numbers[high] {
			return numbers[low]
		}
		mid := low + (high-low)>>1
		if numbers[mid] > numbers[low] {
			low = mid + 1
		} else if numbers[mid] == numbers[low] {
			low++
		} else {
			high = mid
		}
	}
	return numbers[low]
}