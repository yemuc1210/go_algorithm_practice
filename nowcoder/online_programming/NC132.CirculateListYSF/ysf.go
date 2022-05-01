package NC132

/**
 * 环形链表的约瑟夫问题
 * @param n int整型 
 * @param m int整型 
 * @return int整型
*/
/*
编号1-n     m离开
1. 循环
2. 规律
这题和链表没关系   
*/
func ysf( n int ,  m int ) int {
    if n==0 {
		return -1
	}
	ls := make([]int,n+1)
	pos := 1
	for i := 0;i<n-1;i++{
		for j:=0;j<m-1;j++{
			pos = (pos+m-1)%len(ls)
			//删除pos
			
		}
	}
	return ls[0]
}

// #
// # 
// # @param n int整型 
// # @param m int整型 
// # @return int整型
// #
// class Solution:
//     def ysf(self , n , m ):
//         ls=list(range(1,n+1))
//         pos=0
//         for i in range(n-1):
// #             for j in range(m-1):
//             pos=(pos+m-1)%len(ls)
//             del ls[pos]
//         return ls[0]
       
