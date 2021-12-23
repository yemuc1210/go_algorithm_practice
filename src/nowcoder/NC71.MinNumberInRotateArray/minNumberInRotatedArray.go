package NC71

/**
 * 旋转数组的最小数字
 * @param rotateArray int整型一维数组 
 * @return int整型
*/
func minNumberInRotateArray( rotateArray []int ) int {
    // write code here
    	if len(rotateArray) == 0{
		return -1
	}
	if len(rotateArray)==1{
		return rotateArray[0]
	}
    left, mid, right := 0, 0, len(rotateArray) - 1
    for left < right {
        mid = left + (right - left) / 2
        if rotateArray[mid] > rotateArray[right] {
            left = mid + 1
        } else if rotateArray[mid] == rotateArray[right] {
            right = right - 1
        } else {
            right = mid
        }
    }

    return rotateArray[left]
}