package NC63

/**
 * 扑克牌顺子    offer61
 * @param numbers int整型一维数组 
 * @return bool布尔型
*/
func IsContinuous( numbers []int ) bool {
    // A1 J-11   Q-12  K-13   Joker-0-any
	return isStraight1(numbers)
}

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
