package lt852

func peakIndexInMountainArray(arr []int) int {
	//山峰左右大小关系不一致
	//已知长度大于等于3
	if arr[0] > arr[1] {
		return 0
	}
	for i:=1;i<=len(arr)-2;i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		}
	}
	return len(arr)-1
}