package NC104

import (
	"strconv"
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 比较版本号的大小 版本号由修订号组成，修订后之间由.连接 一个修订号可能有多为数字
 * 比较规则：从左到右比较，忽略前导0；如果没有指定位置的修订号，是为0，即1.1=1.1.0.0；
 * @param version1 string字符串
 * @param version2 string字符串
 * @return int整型 1表示v1>v2   -1表示v1<v2
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