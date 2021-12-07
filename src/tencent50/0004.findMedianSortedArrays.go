package tencent50
/* 困难
4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。
*/
/*
时间复杂度要求就决定了需要使用  类似二分搜索的思路
由于要找到最终合并以后数组的中位数，两个数组的总大小也知道，所以中间这个位置也是知道的。
只需要二分搜索一个数组中切分的位置，另一个数组中切分的位置也能得到。
为了使得时间复杂度最小，所以二分搜索两个数组中长度较小的那个数组。

切分方法：nums1[midA-1] ≤ nums2[midB] && nums2[midB-1] ≤ nums1[midA]
如果 `nums1[midA] < nums2[midB-1]`，说明 `midA` 这条线划分出来左边的数小了，切分线应该右移；
如果 `nums1[midA-1] > nums2[midB]`，说明 midA 这条线划分出来左边的数大了，切分线应该左移。


*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
			return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
			// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
			// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
			nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
			nums2Mid = k - nums1Mid
			if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
					high = nums1Mid - 1
			} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
					low = nums1Mid + 1
			} else {
					// 找到合适的划分了，需要输出最终结果了
					// 分为奇数偶数 2 种情况
					break
			}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
			midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
			midLeft = nums1[nums1Mid-1]
	} else {
			midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
			return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
			midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
			midRight = nums1[nums1Mid]
	} else {
			midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a int, b int) int {
	if a > b {
			return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
			return b
	}
	return a
}