package tencent50
/*
88. 合并两个有序数组
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，
另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

 进阶：你可以设计实现一个时间复杂度为 O(m + n) 的算法解决此问题吗？
*/
/*
归并的思路：借助另一个长度为m的数组  【非最优解】
不借助数组，每次插入需要移位  这是从前往后
换个思路，从后往前插入，先选择最后一个元素
	nums1后面n个元素是0，插入不会存在冲突
*/
func merge1(nums1 []int, m int, nums2 []int, n int)  {
	if n==0 {
		return 
	}
	if m==0 {
		copy(nums1,nums2)
		return 
	}
	curIdx1,curIdx2 := m-1,n-1   //两个数组当前元素下标
	for idx:=m+n-1; curIdx1>=0 || curIdx2>=0 ; idx -- {
		var cur int
		//curIdx1 curIdx2 可能为-1  只会有一个
		if curIdx1==-1{
			cur = nums2[curIdx2]
			curIdx2 --
		}else if curIdx2==-1{
			cur = nums1[curIdx1]
			curIdx1--
		}else if nums1[curIdx1] > nums2[curIdx2] {
			cur = nums1[curIdx1]
			curIdx1 -- 
		}else{
			cur = nums2[curIdx2]
			curIdx2--
		}
		nums1[idx] = cur
	}
}