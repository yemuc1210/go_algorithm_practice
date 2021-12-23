package NC20

import "strconv"

/**
 * 复原IP地址
 * 给定只包含数字的字符串 得到所有可能的IP   lt93   offer82-offer II 87
 * @param s string字符串
 * @return string字符串一维数组
 */

/*ly93  回溯
剑指 Offer II 087. 复原 IP
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。

有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。

*/

/*
递归
**切割问题就是利用回溯搜索法把所有可能搜索出来
*/
func restoreIpAddresses(s string) []string {
	var res,path []string
	backTracking(s,path,0,&res)
	return res
}
func backTracking(s string,path []string,startIndex int,res *[]string){
	//终止条件
	if startIndex==len(s)&&len(path)==4{
		tmpIpString:=path[0]+"."+path[1]+"."+path[2]+"."+path[3]
		*res=append(*res,tmpIpString)
	}
	for i:=startIndex;i<len(s);i++{
		//处理
		path:=append(path,s[startIndex:i+1])
		if i-startIndex+1<=3&&len(path)<=4&&isNormalIp(s,startIndex,i){
			//递归
			backTracking(s,path,i+1,res)
		}else {//如果首尾超过了3个，或路径多余4个，或前导为0，或大于255，直接回退
			return
		}
        //回溯
		path=path[:len(path)-1]
	}
}
func isNormalIp(s string,startIndex,end int)bool{
	checkInt,_:=strconv.Atoi(s[startIndex:end+1])
	if end-startIndex+1>1&&s[startIndex]=='0'{//对于前导 0的IP（特别注意s[startIndex]=='0'的判断，不应该写成s[startIndex]==0，因为s截取出来不是数字）
		return false
	}
	if checkInt>255{
		return false
	}
	return true
}