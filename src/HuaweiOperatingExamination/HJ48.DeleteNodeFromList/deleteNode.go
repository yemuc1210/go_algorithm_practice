package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
//     "nc_tools"

)
type ListNode1 struct{
    Val int
    Next *ListNode1
}

//在befor之后插入after值
func insert(head *ListNode1, before,after int) {
	//在此场景下， 不用考虑head==nil
	//p 指向before
	p := head
	for p!=nil &&p.Next!= nil && p.Val!=before{
		p = p.Next
	}
    //可能p.Next=nil
    if p.Next==nil {
        p.Next = &ListNode1{Val:after}
        return 
    }
	//此时p指向before值  在p后插入
	node :=&ListNode1{Val: after, Next: p.Next}
	p.Next = node
}
//删除值为x的节点
func delete(head *ListNode1, x int) {
	p,q := head,head
	for p.Next!=nil && p.Val != x {
		q = p
		p = p.Next
	}
    if p.Next==nil {
        if p.Val == x {
            q.Next = nil
            return 
        }
    }
	//此时p-x  q-p.next
	q.Next = p.Next
    
}

func printList(head *ListNode1) {
	for p:=head; p!=nil ;p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
}
func main(){
    input := bufio.NewReader(os.Stdin)
    data,_ := input.ReadString('\n')
    //根据data 构建链表
    dataSlice := strings.Split(data," ")
	fmt.Println(dataSlice,len(dataSlice),dataSlice[0],dataSlice[len(dataSlice)-1])
	nums := make([]int, len(dataSlice))
	// fmt.Println(nums)
	for i:=0;i<len(dataSlice);i++{
		nums[i],_ = strconv.Atoi(dataSlice[i])
		fmt.Println(nums[i])
	}
	fmt.Println(nums,len(nums))
    dummyHead := &ListNode1{Val:math.MinInt64 ,Next: &ListNode1{Val: nums[1]}}
    
    //剩下两个为一组
    for i:=2;i<len(dataSlice)-1;i+=2{
		insert(dummyHead,nums[i],nums[i+1])
    }
	delete(dummyHead,nums[len(nums)-1])
	//输出
	printList(dummyHead.Next)
}
//6 2 1 2 3 2 5 1 4 5 7 2 2
//5 2 3 2 4 3 5 2 1 4 3
	//2 5 4 1


	// package main

	// import (
	//    "fmt"
	// )
	
	
	
	// func main(){
	
	// 	L := make([]int, 10002)
	// 	n, h,a, b := 0,0,0,0
	// 	fmt.Scanf("%d %d", &n, &h)
	// 	L[h] = 0
	// 	for n = n-1; n >0; n--{
	// 		fmt.Scanf("%d %d", &a, &b)
	// 		L[a] = L[b]
	// 		L[b] = a
	// 	}
	
	// 	// skip or delete
	// 	fmt.Scanf("%d", &a)
	// 	for h > 0{
	// 		if h == a{
	// 			h = L[a]
	// 		}else{
	// 			fmt.Printf("%d ", h)
	// 			h = L[h]
	// 		}
	// 	}
	// }