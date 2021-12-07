package offer61

import (
	"sort"
)

/*
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。
A 不能视为 14。


*/
func isStraight(nums []int) bool {
	if len(nums)!=5{
		return false
	}
	//nums 转化为递增序列
	sort.Ints(nums)
	//记录每个数字之间差值的和
	sub:=0
	for i:=0;i<4;i++{
		if nums[i]==0{
			//continue 不进行下面的代码，进入下一次循环
			continue
		}
		//如果扑克牌（非0数字）重复，不是顺子
		if nums[i]==nums[i+1]{
			return false
		}
		//差值记录
		sub+=nums[i+1]-nums[i]
	}
	//总的差值小于5（或sub<=4）则为顺子
	return sub<5

}

//非排序法
func isStraight1(nums []int) bool {
    //time 记录扑克牌0~13的出现情况
    time := make(map[int]bool,14)
    max,min:=0,14
    zero:=0
    for i:=0;i<5;i++{
     	//如果数字是大小王0
        if nums[i]==0{
            zero++
            //continue 不进行下面的代码，进入下一次循环
            continue
        }
        //如果扑克牌（非0数字）重复，不是顺子
        if time[nums[i]]{
            return false
        }
        max=Max(nums[i],max)
        min=Min(nums[i],min)
        //记录出现过的数字
        time[nums[i]]=true
    }
    //是否存在大小王，返回判断结果
    if zero==0{
        return (max-min)==4
    }
    return max-min<=4
}

//Max 获取相对最大值
func Max(a,b int)int{
    if a>b{
        return a
    }
    return b
}

//Min 获取相对最小值
func Min(a,b int)int{
    if a>b{
        return b
    }
    return a
}
