package main

import (
    "fmt"
    "sort"
)
func main() {
    var num int
    for {
        n,_ := fmt.Scanln(&num)
        if n==0 {
            break
        }
        nums1 := make([]int,num)
        for i:=0;i<num;i++{
            fmt.Scanf("%d%c",&nums1[i])
        }
        fmt.Scanln(&num)
        nums2 := make([]int,num)
        for i:=0;i<num;i++{
            fmt.Scanf("%d%c",&nums2[i])
        }
        //排序
        sort.Slice(nums1, func(i,j int)bool{
            return nums1[i] < nums1[j]
        })
        sort.Slice(nums2, func(i,j int)bool{
            return nums2[i] < nums2[j]
        })
        // 各自去重
        var last int = 0
        for i:=1;i<len(nums1);{
            for i<len(nums1) && nums1[i] == nums1[last]{
                i ++
            }
            if i<len(nums1){
                // 迁移
                nums1[last+1] = nums1[i]
                last = last+1
            }
        }
		// fmt.Println(nums1)
        nums1 = nums1[:last+1]
		// fmt.Println(nums1)
		last=0
        for i:=1;i<len(nums2);{
            for i<len(nums2) && nums2[i] == nums2[last]{
                i++
            }
            if i<len(nums2){
                // 迁移
                nums2[last+1] = nums2[i]
                last = last+1
            }
        }
		nums2 = nums2[:last+1]
        //合并
        var i,j int
        var res []int
        for i<len(nums1) && j<len(nums2) {
            if nums1[i] < nums2[j] {
                res = append(res,nums1[i])
                i++
            }else if nums1[i] > nums2[j] {
                res = append(res, nums2[j])
                j++
            }else{
                res = append(res, nums1[i])
                i++
                j++
            }
        }
        // 判断哪个数组还有数字
        if len(nums1) != 0 {
            var k int = len(res)-1
            for i < len(nums1) {
                if res[k] != nums1[i] {
                    res = append(res, nums1[i])
                    i++
                    k++
                }
            }
        }
        if len(nums2) != 0 {
            var k int = len(res)-1
            for j < len(nums2) {
                if res[k] != nums2[j] {
                    res = append(res, nums2[j])
                    j++
                    k++
                }
            }
        }
        //out
        for m:=0;m<len(res);m++{
            fmt.Print(res[m])
        }
        fmt.Println()
    }
}