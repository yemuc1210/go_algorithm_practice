package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main() {
    var v1,v2 string
    fmt.Scanf("%s %s",&v1, &v2)
    
    //不用判断字符格式的合法性
    
    //转换成数字比较大小
    res := compareVersion(v1,v2)
	fmt.Println(res)
}

func compareVersion(version1,version2 string) int {
    v1 := strings.Split(version1,".")
    v2 := strings.Split(version2,".")
    
    //
	for i:=0;i<len(v1) || i<len(v2);i++{
		//对于每个子版本逐个比较
		var x,y int
		if i<len(v1) {
			x,_  = strconv.Atoi(v1[i])

		}
		if i<len(v2) {
			y,_ = strconv.Atoi(v2[i])
		 
		}
		if x>y {
			return 1
		}else if x<y {
			return -1
		}else{
			continue
		}
	}
	return 0
}