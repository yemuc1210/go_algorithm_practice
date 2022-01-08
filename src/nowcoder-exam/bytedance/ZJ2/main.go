package main

import "fmt"
// 比赛场次能被3整除且当前球队最高分小于等于n/3，则有可能踢平，步骤：
// 1、n是否能被3整除，否输出no，是进行2
// 2、求满足(k,d1,d2)的所有最大得分的最小值
// 3、如果这个最小值≤n/3，输出yes，否则输出no
func main() {
	var t int
	fmt.Scanln(&t)
	var tmp, left int
	for i := 0; i < t; i++ {
		var n, k, d1, d2 int
		fmt.Scanln(&n, &k, &d1, &d2)
		if n%3 != 0 {
			fmt.Println("no")
			continue
		}
		//1<2<3   3m = k-d1-d2-d2
		tmp = k - d1 - d1 - d2
		if tmp >= 0 && tmp%3 == 0 {
			left = (n - k) - (d1 + d2 + d2) //剩余-need
			if left >= 0 && left%3 == 0 {
				fmt.Println("yes")
				continue

			}
		}
		//1<2  2>3   m m+d1 m+d1-d2   = k
		tmp = k - d1 - d1 + d2
		if tmp >= 0 && tmp%3 == 0 {
			//打平需要d1+d2   1、3追上2
			left = (n - k) - (d1 + d2)
			if left >= 0 && left%3 == 0 {
				fmt.Println("yes")
				continue
			}
		}
		//1>2>3   m m-d1  m-d1-d2 = k
		tmp = k + d1 + d1 + d2 //3m
		if tmp >= 0 && tmp%3 == 0 {
			//2 3 追1
			left = (n - k) - (d1 + d1 + d2)
			if left >= 0 && left%3 == 0 {
				fmt.Println("yes")
				continue

			}
		}

		//1>2  2<3  m m-d1 m-d1+d2
		tmp = k + d1 + d1 - d2
		if tmp >= 0 && tmp%3 == 0 {
			left = (n - k) - (d1 + d1 - d2)
			//需要区分d1 d2的大小
			if d1 >= d2 {
				//1 得分最多
				// 2 3 追1
				left = (n - k) - (d1 + d1 - d2)
				if left >= 0 && left%3 == 0 {
					fmt.Println("yes")
					continue
				}

			} else {
				//3得分最多
				left = (n - k) - (d2 + d2 - d1)
				if left >= 0 && left%3 == 0 {
					fmt.Println("yes")
					continue
				}
			}
		}
		//否则不在合法情况
		fmt.Println("no")
	}

}
