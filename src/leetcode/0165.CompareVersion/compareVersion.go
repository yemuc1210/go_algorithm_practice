package lt165

import (
	"strconv"
	"strings"
)

/*
每日一题 0901
给你两个版本号 version1 和 version2 ，请你比较它们。

版本号由一个或多个修订号组成，各修订号由一个 '.' 连接。每个修订号由 多位数字 组成，
可能包含 前导零 。每个版本号至少包含一个字符。
修订号从左到右编号，下标从 0 开始，最左边的修订号下标为 0 ，下一个修订号下标为 1 ，
以此类推。例如，2.5.33 和 0.1 都是有效的版本号。

比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较 忽略任何前导零后的整数值 。
也就是说，修订号 1 和修订号 001 相等 。
如果版本号没有指定某个下标处的修订号，则该修订号视为 0 。
例如，版本 1.0 小于版本 1.1 ，因为它们下标为 0 的修订号相同，
	而下标为 1 的修订号分别为 0 和 1 ，0 < 1 。

返回规则如下：

	如果 version1 > version2 返回 1，
	如果 version1 < version2 返回 -1，
	除此之外返回 0。


*/
//字符串分割  然后分割的数字转化成整数，可能存在超大版本号，但实际中不会出现这种
func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1,".")
	v2 := strings.Split(version2,".")

	//得到的是字符串数组
	for i:=0;i<len(v1) || i<len(v2);i++ {
		x,y := 0,0
		if i < len(v1) {
			x,_  = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			y,_ = strconv.Atoi(v2[i])
		}
		if x>y {
			return 1
		}
		if x<y{
			return -1
		}
	}
	return 0
}