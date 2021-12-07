/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
package leetcode

import "structs"
type ListNode = structs.ListNode

//需要记录进位carry   
func plusOne(digits []int) []int{
	carry,index := 0,len(digits)-1
	carry = (digits[index]+1)/10
	digits[index] = (digits[index] + 1)%10
	index --
	/*
	注意，[9]->[1,0]
	*/
	for carry == 1 && index >=0{   
		// plusOne(digits[:index-1])//递归不好，空间消耗大
		carry = (digits[index]+1)/10
		digits[index] = (digits[index] +1)%10
		index --
	}
	if carry==1 && index==-1{//最高位有进位，在数组最前面加一位，感觉链表好一些
		digits = append(digits, 1) //append函数会自动扩容
		for i:=len(digits)-1;i>0;i--{
			digits[i] = digits[i-1]
		}
		digits[0] = 1
	}
	return digits
}

//使用链表再试一次
//算了，也不是很好，还要浪费一份空间制造链表
// func plusOne_ListNode(digits []int)[]int{
// 	carry := (digits[len(digits)-1] +1) /10
// 	rear := &ListNode{Val: digits[len(digits)-1] }
// 	head := rear
// 	for 




// }