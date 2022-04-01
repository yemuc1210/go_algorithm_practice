package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	// 自身可以和每个数位都整除  所以不能包括0（不能做被除数）
	// 返回[left,right]之间的自除数
	var res []int
	for left <= right {
		if left%10 == 0 {
			left++
			continue
		}
		tmp := left

		for tmp != 0 {
			// 取最后一位
			//num := tmp % 10
			if tmp%10 == 0 {
				break
			}
			if left%(tmp%10) == 0 {
				tmp /= 10
			}else{
				// goto out
				break
			}
			
			// 否则当前left不是自除数
		}

		if tmp == 0 {
			res = append(res, left)
		}
		left++
	}
	return res
}

func main() {
	fmt.Println("vim-go")
	res := selfDividingNumbers(1, 22)
	fmt.Println(res)
}
