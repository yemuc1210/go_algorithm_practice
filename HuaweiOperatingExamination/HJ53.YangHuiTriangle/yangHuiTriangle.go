// package main

// import "fmt"

// func main() {
// 	var N int
// 	var n int
// 	var err error
// 	for { //对应多组输入
// 		n,err = fmt.Scanln(&N)
// 		if n==0 || err!=nil {
// 			break
// 		}
// 		//输出第N行第一个偶数出现的位置  1<=N<=10^9
// 		// 需要两个数组，记录两行的数
// 		if N<=2 {
// 			fmt.Println(-1)  //第1 2行没有偶数
// 		}
// 		//迭代计算到第n行
// 		arr1,arr2 := []int{1,1,1},[]int{1}   //使用切片，可以自动扩容
// 		for i:=3;i<=N;i++{ //从第三行开始
// 			//计算第三行数字
// 			//第i行有 2(n-1)+1 个数
// 			//一种方式是利用go切片性质 对arr1扩容 
// 			for j:=1;j<2*(i-1)+1;j++{
// 				//第一个数是1，不必说
// 				// arr2[j] = arr1[j-1] + arr1[j]+ arr[j-2]   
// 				//按照这个公式  会有越界的情况
// 				if j==1 {
// 					arr2 = append(arr2, arr1[j-1] + arr1[j])   
// 				}else if j==2*(i-1)-1 {
// 					//倒数第二个数  如i=3  j=3
// 					arr2 = append(arr2,arr1[j-2] + arr1[j-1]) 
// 				}else if j==2*(i-1) {
// 					//最后一个数  i=3 j=4
// 					arr2 = append(arr2,1)
// 				}else{
// 					//中间的根据公式正常计算 i==3 j==2 
// 					arr2 = append(arr2,arr1[j-2] + arr1[j-1] + arr1[j])
// 				}
// 			}
// 			//如果当前是第N行 ，计算结果
// 			if i == N {
// 				//遍历一次
// 				flag := false
// 				for i:=0;i<len(arr2);i++{
// 					//找偶数
// 					if arr2[i]%2==0 {
// 						fmt.Println(i+1)
// 						//跳出
// 						flag = true
// 						break
// 					}
// 				}
// 				if !flag {
// 					fmt.Println(-1)
// 				}
// 			}
// 			// 交换两个切片
// 			// copy(arr1,arr2)  //将arr2copy到arr1
// 			arr1 = arr2    //直接交换引用
// 			arr2 = []int{1}  //初始化
// 		}
// 	}
// }

// //会超时  计算每个一半即可
// 根据规律  
package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
)

func main(){
    var params []string
    input := bufio.NewScanner(os.Stdin)
    for input.Scan(){
        params = append(params, input.Text())
    }
    for loop:= 0; loop< len(params); loop++{
        res := OutPrint(params[loop])
        fmt.Println(res)
    }
}

func OutPrint(s string) int {
    var res int
    num, _ := strconv.Atoi(s)
    if num ==1 || num ==2 { 
        res = -1
    } else if num%4 ==1 || num%4 ==3 {   //
        res = 2
    }else if num%4 ==0 {  
        res = 3
    }else if num%4 ==2 { 
        res = 4
    }
    return res
}