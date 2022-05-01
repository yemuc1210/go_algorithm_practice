package lt1984

import (
	// "crypto/aes"
	// "fmt"
	"math"
	"sort"
)
func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)
    ans := math.MaxInt32
    for i, num := range nums[:len(nums)-k+1] {
        ans = min(ans, nums[i+k-1]-num)
    }
    return ans
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

// 快排
func quickSort(nums []int, left,right int) []int {
    if left <= right {
        return nums
    }
    x := nums[(left+(right-left)/2)]
    i,j := left-1,right+1
    for i<j {
        for {
            i++
            if nums[i] >= x {
                break
            }//找到左边大于
        }
        for {
            j--
            if nums[j] <= x {
                break
            }
        }
        // 交换
        if i<j {
            nums[i],nums[j] = nums[j],nums[i]
        }
    }
    quickSort(nums,left,j)
    quickSort(nums,j+1, right)
	return nums
}