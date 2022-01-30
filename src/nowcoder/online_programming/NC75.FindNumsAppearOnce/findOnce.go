package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型一维数组
 * @return int整型一维数组
 */
func FindNumsAppearOnce( nums []int ) []int {
    xorSum := 0
    for _, num := range nums {
        xorSum ^= num   //两个出现一次的数的异或和
    }
    lsb := xorSum & -xorSum   //得到最低位的1
	// fmt.Printf("lsb=%b\n",lsb)   // lsb = 10    xorSum = 6 = Ox 0110
    type1, type2 := 0, 0
    for _, num := range nums {  //直接根据最终的结果
        if num&lsb > 0 {  //一类
            type1 ^= num
        } else {
            type2 ^= num
        }
    }
	if type1 > type2 {
		return []int{type2,type1}
	}
    return []int{type1, type2}
}

func main(){
	input := []int{1,4,1,6}  //4^6=2   100 110  011  001
	res := FindNumsAppearOnce(input)
	fmt.Println(res)
}