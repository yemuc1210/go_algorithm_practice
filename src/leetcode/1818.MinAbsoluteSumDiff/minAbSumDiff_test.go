package lt1818

import (
	"fmt"
	"testing"
)

type question1818 struct{
	nums1,nums2 []int
	ans int
}

func Test_Pro1818(t *testing.T){
	qs:=[]question1818{
		{
			nums1: []int{1,7,5},
			nums2: []int{2,3,5},
			ans: 3,
		},
		{
			nums1: []int{2,4,6,8,10},
			nums2: []int{2,4,6,8,10},
			ans: 0,
		},
		{
			nums1: []int{1,10,4,4,2,7},
			nums2: []int{9,3,5,1,7,4},
			ans: 20,
		},
		{
			nums1: []int{1,28,21},
			nums2: []int{9,21,20},
			ans: 9,
		},
	}
	for _,ex:=range qs{
		nums1,nums2,as := ex.nums1,ex.nums2,ex.ans
		res:=minAbsoluteSumDiff(nums1,nums2)
		if res!= as{
			t.Errorf("error [input]:%v %v [out]:%v   [ans]:%v\n",nums1,nums2,res,as)
		}else{
			fmt.Println("pass")
		}
	}
}