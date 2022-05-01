package main

import (
	// "bytes"
	"fmt"
)

func main() {
    var input string
    var n int
    var err error
    for {
        n,err = fmt.Scanln(&input)
        if n==0 {
            break
        }
        if err!=nil {
            panic(err)
        }
        
        //按照ASCII值排序
        bytes := []byte(input)
        
        quickSort(&bytes,0,len(bytes)-1)
		fmt.Println(bytes)
    }
}

//快排
func quickSort(bytes *[]byte,l,r int) {
	// l,r := 0,len(*bytes)-1
	for l < r{
		idx := partition(bytes,l,r)
		quickSort(bytes,l,idx-1)
		quickSort(bytes,idx+1,r)
	}
}

//返回pivot下标
func partition(bytes *[]byte, l,r int)  int {
	left,right := l,r
	pivot := (*bytes)[left] 
	//区间[left,right]
	for left<right {
		//右边找到小于等于pivot
		for left<right && (*bytes)[right] > pivot {
			right --
		}
		(*bytes)[left] = (*bytes)[right]
		//左边找到大于等于pivot的
		for left<right && (*bytes)[left] < pivot{
			left ++
		}
		(*bytes)[right] = (*bytes)[left]
	}	
	//交换pivot和left
	(*bytes)[left] = pivot
	return left
}
