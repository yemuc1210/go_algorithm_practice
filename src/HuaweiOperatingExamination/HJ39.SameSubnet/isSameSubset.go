package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
    var mask string
    var ip1,ip2 string
    for {
        //输入  一组三行
        n,_ := fmt.Scanln(&mask)
        if n== 0 {
            break
        }
        //输入ip
        fmt.Scanln(&ip1)
        fmt.Scanln(&ip2)
        //处理数据
        // 四段
        masks := strings.Split(mask,".")
        ip1s := strings.Split(ip1,".")
        ip2s := strings.Split(ip2,".")
		// 判断格式是否合法
		if !isValisMask(masks) || !isValidIP(ip1s) ||!isValidIP(ip2s) {
			// fmt.Println(isValisMask(masks))
			// fmt.Println(isValidIP(ip1s))
			// fmt.Println(isValidIP(ip2s))
			fmt.Println(1)
			continue
		}
		r1 := solve(masks,ip1s)
		r2:= solve(masks,ip2s)

		if isEqual(r1,r2) {
			fmt.Println(0)
		}else{
			fmt.Println(2)
		}
			
	}
}

func solve(masks, ips []string)(res  []int) {
	// 进行与运算   255.255.255.0
// 192.168.224.256
// 192.168.10.4
	for i:=0;i<4;i++{
		mask,_ := strconv.Atoi(masks[i])
		sub,_ :=  strconv.Atoi(ips[i]) 
		// 两个数字 与 运算	
		res = append(res, sub&mask)  //直接与
	}
	return 
}
func isEqual(a,b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i:=0;i<len(a);i++{
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func isValidIP(ips []string) bool {
	if len(ips) != 4 {
		return false
	}
	//ip位置必须是0-255
	for i:=0;i<len(ips);i++{
		ip,_ := strconv.Atoi(ips[i])
		if ip<0 || ip>255 {
			return false
		}
	}
	return true
}
func isValisMask(masks []string) bool {
	// 掩码不能全部为1  前面为1 后面为0
	if masks[3] == "255" {
		return false  // 不管怎么样都是非法的
	}
	flag := true
	
	for i:=0;i<=3;i++ {
		if masks[i] == "255" && flag{
			continue
		}else{
			//那么后面都不能是255 且必须在0-254之间
			if masks[i] == "0" {
				//那么后面必须是0
				// fmt.Println(masks[i])
				for j:=i+1;j<=3;j++{
					// fmt.Println(masks[j])
					if masks[j] != "0" {
						return false
					}
				}
			}
			flag = false
			mask,_ := strconv.Atoi(masks[i])
			if mask<0 || mask>254 {
				return false
			}else{
				continue
			}
		}		
	}
	// 如果之前没有return 
	return true
}