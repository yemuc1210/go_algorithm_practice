package lt458

import "fmt"

/*
458. 可怜的小猪
有 buckets 桶液体，其中 正好 有一桶含有毒药，其余装的都是水。它们从外观看起来都一样。为了弄清楚哪只水桶含有毒药，你可以喂一些猪喝，通过观察猪是否会死进行判断。不幸的是，你只有 minutesToTest 分钟时间来确定哪桶液体是有毒的。

喂猪的规则如下：

选择若干活猪进行喂养
可以允许小猪同时饮用任意数量的桶中的水，并且该过程不需要时间。
小猪喝完水后，必须有 minutesToDie 分钟的冷却时间。在这段时间里，你只能观察，而不允许继续喂猪。
过了 minutesToDie 分钟后，所有喝到毒药的猪都会死去，其他所有猪都会活下来。
重复这一过程，直到时间用完。
给你桶的数目 buckets ，minutesToDie 和 minutesToTest ，返回在规定时间内判断哪个桶有毒所需的 最小 猪数。
*/

/*
1. 第一轮
	拿一只小猪同时喝n/2桶水  二分   死一只  另一只 minutesTest后可以继续使用
	minuteDie后
2. 第二轮     问题是第二轮何时开始： 需要比较Die和Test的大小关系
	再次二分
	若 Test <= Die 则第一轮的小猪可以继续使用
		到这一轮结束，共使用3只
	若 Test > Die 则出现测试结果时，第一轮存活小猪的CD还没到
		那这一轮结束，共使用4只

问题关键，比较Test Die直接的大小（倍数）关系
举例
	1只猪 15min(Die)  1h 最多检验 5个   前四次没死，第五次死
	2只猪 可以让其中1只猪一次喝5桶水（行 1h行 5次25桶）  另一只和列5桶      最多检验25桶
	3只  5*5*5 = 125
*/
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	pigs := 0 
	for myPow((minutesToTest/minutesToDie+1),pigs) < buckets {
		pigs ++
	}
	return pigs
}

func myPow(a,b int) int{
	fmt.Println(a,b)
	//a^b
	res := 1
	for i:=0;i<b;i++ {
		res *= a
		fmt.Println(a)
	}
	return res
}