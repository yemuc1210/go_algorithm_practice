package main

import "fmt"

/*
1247. 交换字符使得字符串相同
有两个长度相同的字符串 s1 和 s2，且它们其中 只含有 字符 "x" 和 "y"，
你需要通过「交换字符」的方式使这两个字符串相同。

每次「交换字符」的时候，你都可以在两个字符串中各选一个字符进行交换。

交换只能发生在两个不同的字符串之间，绝对不能发生在同一个字符串内部。
也就是说，我们可以交换 s1[i] 和 s2[j]，但不能交换 s1[i] 和 s1[j]。

最后，请你返回使 s1 和 s2 相同的最小交换次数，如果没有方法能够使得这两个字符串相同，则返回 -1 。

*/

//     最后一句的解释，x+y为奇数，说明不可能
//     1.若x和y都为偶数
//     例如：
//     XX YY YY
//     YY XX XX 就会进行(x+y)/2次交换，如示例一
//     +1 +1 +1

//     2.若x和y都为奇数
//     XX X YYYY Y
//     YY Y XXXX X 就会进行(x-1)/2+(y-1)/2+2=(x+y)/2+1
//        |______|
//          相连

//     二者均可写作(x+1)/2+(y+1)/2;

func minimumSwap(s1 string, s2 string) int {
    x,y,n := 0,0,len(s1)
    for i:=0;i<n;i++{
		// 只有字符不同得位置需要被交换
        if s1[i] != s2[i] {
            if s1[i]=='x' {
                x++
            }else{
                y++
            }
        }
    }
	// 如果x y 个数有一个是奇数,都无法完成交换
    if (x+y)%2==1 {
        return -1
    }
	// 字符不同  两种   s1[i]==x&&s2[i]==y --> xy   另一种yx
	// 一对xy  一次交换
	// 一个xy 和一个yx交换,需要两次
	// 贪心,让更多得xy yx彼此成对内部消化,剩余得在互相匹配
	// 最后至多剩余一个
	
    return (x+1)/2+(y+1)/2
}

func main() {
	s1 := "xx"
	s2 := "yy"
	res := minimumSwap(s1,s2)
	fmt.Println(res)
}