package lt1414


func findMinFibonacciNumbers(k int) int {
	var table []int=[]int{1,1}
	pre:=1
	last:=1
	for last<k{
		table=append(table,pre+last)
		pre,last=last,pre+last
	}
	num:=0
	right:=len(table)-1
	left:=0
	for k!=0{
		left=0
		for left<right{
			mid:=(left+right+1)>>1
			if table[mid]<=k{
				left=mid
			}else{
				right=mid-1
			}
		}
		num++
		k-=table[left]
	}
	return num
}

func findMinFibonacciNumbers1(k int) int {
    // 首先产生k个数  答案在k个数的范围之内
    nums := make([]int,k+1)
    i := fib(k,&nums)
    // 在nums里找可行解
    nums = nums[0:i]
    // 问题变为找和为k  现在数组是递增有序的
    ans := find(k, nums)
    return ans
}

func find(k int, nums []int) int {
    if k==0 {
        return 0
    }
    l,r := 0,len(nums)-1
    // 二分
    for l < r{
        mid := l + ((r-l)>>1) + 1
        if nums[mid] > k {
            r = mid - 1
        }else {
            l = mid
        }
    }
    return find(k-nums[l],nums) + 1
}

// 产生K个 或 到k为止
func fib(k int, nums *[]int) int {
    (*nums)[0] = 1
    (*nums)[1] = 1
    if k<=1 {
        return k
    }

    var i int

    for i=2;(*nums)[i] <= k;i++ {
        (*nums)[i] = (*nums)[i-1] + (*nums)[i-2]
    }
    return i 
}