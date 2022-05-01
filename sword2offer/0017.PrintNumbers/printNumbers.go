package offer17

import "strconv"

/*
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。
比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]


应该有限制，不然没法返回整数数组
*/

// func printNumbers(n int) []int {
// 	if n==0{
// 		return []int{}
// 	}
// 	res:=[]int{}
// 	max:=0
// 	for n>0 {
// 		max = max*10+9
// 		n--
// 	}
// 	for i:=1;i<=max;i++{
// 		res = append(res, i)
// 	}
// 	return res
// }

//解决大数溢出等异常
//基于递归生成全排列
func printNumbers(n int)[]int{
	res := []int{}

	if n <= 0 {
		return res
	}

	number := make([]byte, n)

	for i := 0; i < n; i++ {
		number[i] = '0'
	}

	for !increment(&number, n) {
		if number[0] == '0' {
			tmp := make([]byte, n)
			copy(tmp, number)
			tmp = tmp[1:]
			add, _ := strconv.Atoi(string(tmp))
			res = append(res, add)
		} else {
			add, _ := strconv.Atoi(string(number))
			res = append(res, add)
		}
	}
	return res
}

func increment(number *[]byte, length int) bool {
	// 只有最高位进位的时候才会超过最大值
	isOverflow := false

	// 进位
	var takeOver byte
	takeOver = 0

	for i := length - 1; i >= 0; i-- {
		nSum := ((*number)[i] - '0') + takeOver

		// 从低位开始+1
		if i == length-1 {
			nSum++
		}

		if nSum >= 10 {
			// 需要进位

			if i == 0 {
				// 最高位不能进位， 返回进位结束
				isOverflow = true
			} else {
				nSum -= 10

				takeOver = 1
				(*number)[i] = '0' + nSum
			}
		} else {
			// 执行一次+1操作就退出
			(*number)[i] = '0' + nSum
			break
		}
	}
	return isOverflow
}
