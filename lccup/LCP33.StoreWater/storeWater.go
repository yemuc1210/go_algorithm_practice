package lcp33

import "math"

/*
bucket  容量
vat  缸的最低蓄水量
操作：
	升级容量
	蓄水

	返回至少需要几次可以完成所有蓄水要求
首先增加容量
倒水
*/
func storeWater(bucket []int, vat []int) int {
	maxVat := 0
	for _,v := range vat {
		maxVat = max(v,maxVat)
	}

	if maxVat == 0 {
		//说明不需要蓄水
		return 0
	}
	res := math.MaxInt64
	for pour:=1; pour<=10000;pour ++{
		if pour >= res{
			break
		}
		upgrade := pour
		for i:=0;i<len(vat);i++{
			//枚举每个水桶，计算总升级次数
			cur := (vat[i]+pour-1)/pour - bucket[i]   ////容量/倒水次数-初始蓄水量=升级次数
			if cur>0{
				upgrade+=cur

			}else{
				upgrade += 0
			}
			if upgrade >= res{
				break
			}
		}
		res = min(res, upgrade+pour)  //总次数=倒水次数+总升级次数
	}
	return res
}
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}