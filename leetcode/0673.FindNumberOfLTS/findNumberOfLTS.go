package lt673
/*
算法计划  300  DP

给定一个未排序的整数数组，找到最长递增子序列的个数。
算法：

假设对于以 nums[i] 结尾的序列，我们知道最长序列的长度 length[i]，以及具有该长度的序列的 count[i]。
对于每一个 i<j 和一个 A[i]<A[j]，我们可以将一个 A[j] 附加到以 A[i] 结尾的最长子序列上。
如果这些序列比 length[j] 长，那么我们就知道我们有count[i] 个长度为 length 的序列。
如果这些序列的长度与 length[j] 相等，
那么我们就知道现在有 count[i] 个额外的序列（即 count[j]+=count[i]）。


*/
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	lengths := make([]int,n)           //以nums[i]结尾的最长序列的长度
	counts := make([]int,n)   //既有该长度的序列
	//初始化
	for i := 0; i < n; i++ {
		lengths[i] = 1
		counts[i] = 1
	}
	max:=1
	res:=0
	for i:=1;i<n;i++{
		for j:=0;j<i;j++{
			if nums[j]<nums[i]{
				if lengths[i] < lengths[j]+1{
					lengths[i] = lengths[j]+1
					counts[i]=counts[j]
				}else if lengths[i]==lengths[j]+1{
					counts[i]+=counts[j]
				}
			}
		}
		if lengths[i]>max{
			max = lengths[i]
		}
	}
	   //然后将结果加上相应的组合长度
	   for i:=0;i<len(counts);i++{
        if lengths[i]==max{
            res+=counts[i]
        }
    }
    return res
}