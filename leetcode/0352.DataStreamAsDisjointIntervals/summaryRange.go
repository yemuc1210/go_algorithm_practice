package lt352
/*
352. 将数据流变为多个不相交区间
 给你一个由非负整数 a1, a2, ..., an 组成的数据流输入，请你将到目前为止看到的数字总结为不相交的区间列表。

实现 SummaryRanges 类：

SummaryRanges() 使用一个空数据流初始化对象。
void addNum(int val) 向数据流中加入整数 val 。
int[][] getIntervals() 以不相交区间 [starti, endi] 的列表形式返回对数据流中整数的总结。
 

*/
// type SummaryRanges struct {
//     *redblacktree.Tree
// }

// func Constructor() SummaryRanges {
//     return SummaryRanges{redblacktree.NewWithIntComparator()}
// }


// /*
// (1) val ∈ [l,r]   无变化
// (2)右边界r紧贴val,r+1=val,则加入val后，区间变为[l,r+1]
// (3)l-1=val   [l-1,r]
// (4)同时满足  [l0,r0] r0+1=val,   [l1,r1] l1-1=val  则合并[l0,r1]
// (5)else    val单独形成 [val,val]
// */
// func (this *SummaryRanges) AddNum(val int)  {
// 	// 找到 l0 最大的且满足 l0 <= val 的区间 interval0 = [l0, r0]
// 	interval0, has0 := ranges.Floor(val)
// 	if has0 && val <= interval0.Value.(int) {
// 		// 情况一
// 		return
// 	}

// 	// 找到 l1 最小的且满足 l1 > val 的区间 interval1 = [l1, r1]
// 	// 在有序集合中，interval1 就是 interval0 的后一个区间
// 	interval1 := ranges.Iterator()
// 	if has0 {
// 		interval1 = ranges.IteratorAt(interval0)
// 	}
// 	has1 := interval1.Next()

// 	leftAside := has0 && interval0.Value.(int)+1 == val
// 	rightAside := has1 && interval1.Key().(int)-1 == val
// 	if leftAside && rightAside {
// 		// 情况四
// 		interval0.Value = interval1.Value().(int)
// 		ranges.Remove(interval1.Key())
// 	} else if leftAside {
// 		// 情况二
// 		interval0.Value = val
// 	} else if rightAside {
// 		// 情况三
// 		right := interval1.Value().(int)
// 		ranges.Remove(interval1.Key())
// 		ranges.Put(val, right)
// 	} else {
// 		// 情况五
// 		ranges.Put(val, val)
// 	}

// }


// func (this *SummaryRanges) GetIntervals() [][]int {
// 	ans := make([][]int, 0, ranges.Size())
//     for it := ranges.Iterator(); it.Next(); {
//         ans = append(ans, []int{it.Key().(int), it.Value().(int)})
//     }
//     return ans

// }


// /**
//  * Your SummaryRanges object will be instantiated and called as such:
//  * obj := Constructor();
//  * obj.AddNum(val);
//  * param_2 := obj.GetIntervals();
//  */

// 主要就是二分查找, 再根据插入的具体位置进行讨论, 看看能否变更现有区间的范围, 合并区间等
// 对于Go语言来说, 要注意使用append来实现slice的插入和删除

type SummaryRanges struct {
    record [][]int // 记录区间范围
}


/** Initialize your data structure here. */
func Constructor() SummaryRanges {
    return SummaryRanges{}
}


func (this *SummaryRanges) AddNum(val int)  { // 左闭右开的二分查找, 寻找lowerbound, 即第一个左边界不小于val的位置
    var sort = func(l, r, val int) int{
        for l<r{
            mid := l + (r-l)/2
            if this.record[mid][0]>=val {
                r = mid
            }else{
                l= mid + 1
            }
        }
        return l
    }  

    l, r := 0, len(this.record)
     if r==0{ // 为空的情况
        this.record = append(this.record, []int{val, val})
        return
    }  
    loc := sort(l, r, val)
   
    if loc == 0 { // 插在最左边
        if this.record[0][0] -1 == val { // 可以和最左边的左边界连起来
            this.record[0][0]=val
            return 
        }
        if this.record[0][0] -1 > val{ // 不可以和最左边的左边界连起来
            this.record = append([][]int{{val, val}}, this.record...)
            return 
        }
        return 
    } 

    if loc == r { // 插在最右边
        if this.record[r-1][1] +1 == val { // 可以和最右边的右边界连起来
            this.record[r-1][1]=val
            return 
        }
        if this.record[r-1][1]+1< val { // 不可以和最右边的右边界连起来
            this.record = append(this.record, []int{val, val})
            return 
        }
        return 
    }

    // 插在中间的四种情况
    if this.record[loc-1][1] +1== val && this.record[loc][0]-1 == val { // 可以同时和左右连起来
        this.record[loc-1][1]=this.record[loc][1]
        this.record=append(this.record[0:loc], this.record[loc+1:]...)
        return
    }
    if this.record[loc-1][1] +1== val { // 只能和左边连起来
        this.record[loc-1][1]=val
        return
    }
    if this.record[loc][0]-1 == val { // 只能和右边连起来
        this.record[loc][0]= val
        return
    }
    if val < this.record[loc][0] && val>this.record[loc-1][1]{ // 左右都连不起来, 只能插在中间
        tmp := append([][]int{},this.record[loc:]...)
        this.record = append(this.record[0:loc], []int{val, val})
        this.record = append(this.record, tmp...)
        return
    }

}



func (this *SummaryRanges) GetIntervals() [][]int {
    var res [][]int 
    for _, val := range this.record {
        res = append(res, val)
    }
    return res
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */

