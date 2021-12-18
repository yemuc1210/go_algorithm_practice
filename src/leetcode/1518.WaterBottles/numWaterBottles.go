package lt1518

func numWaterBottles(numBottles int,  numExchange int) int {
	// n+(n-1)/(m-1) m>1
	return numBottles + (numBottles-1)/(numExchange-1)
}

/*
等价交换
拿n=9, m=3这个为例。
3个空瓶=1个啤酒=1个空瓶+1单位酒 --> 2个空瓶=1单位酒 --> 1个空瓶=0.5单位酒
那么该人总共可以喝 公式A:n+n/(m-1)=9+9/(3-1)=13瓶，但是这个答案是不完美的。
// n+(n-1)/(m-1) m>1    所以我们要对n-1 减去的1就是最后一次喝的用的空瓶
*/