package lt108
import "structs"
type TreeNode = structs.TreeNode
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
		return nil
	}
	return &TreeNode{Val: nums[len(nums)/2], Left: sortedArrayToBST(nums[:len(nums)/2]), Right: sortedArrayToBST(nums[len(nums)/2+1:])}
}