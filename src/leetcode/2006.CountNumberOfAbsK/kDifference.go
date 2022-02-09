package lt2006

func countKDifference(nums []int, k int) int {
    var res int
    for i:=0;i<len(nums);i++{
        for j:=i+1;j<len(nums);j++{
            if abs(nums[i],nums[j]) == k {
                res++
            }
        }
    }
    return res
}
func abs(a,b int) int {
    if a>b {
        return a-b
    }else{
        return b-a
    }
}