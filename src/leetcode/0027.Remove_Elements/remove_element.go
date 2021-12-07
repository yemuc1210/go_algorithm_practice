/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，
并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/

package leetcode


//原地移除，又是移位，
//不过函数只要返回长度值，用一个标记数组   不可行，不允许使用额外数组空间
func removeElement(nums []int, val int) int {
	//方法是先计算出有几个重复的，然后集中移位
	/*移位方法是从前往后
		遇到第一个，前移一位直到遇见下一个目标
		遇到第二个，前移两位直到遇见下一个目标
		...
	*/
	steps := 0 // 当前需要的移动步数
	for i:=0;i<len(nums); {
		if nums[i] == val {//需要移位
			steps ++
			//下面移位
			j:=i+1
			for j<len(nums)&& nums[j]!=val{//移位，直到遇到下一个val
				nums[j-steps] = nums[j]   //前移动steps
				j++
			}//跳出循环时，num[j]==val   得到j的位置
			//下一个循环从第二个j之后开始啊
			i = j
		}else{
			i ++ 
		}
	}

	return len(nums)-steps
}
