package lt468

import (
	"fmt"
	"strconv"
	"strings"
)

/*
468. 验证IP地址
编写一个函数来验证输入的字符串是否是有效的 IPv4 或 IPv6 地址。

如果是有效的 IPv4 地址，返回 "IPv4" ；
如果是有效的 IPv6 地址，返回 "IPv6" ；
如果不是上述类型的 IP 地址，返回 "Neither" 。
IPv4 地址由十进制数和点来表示，每个地址包含 4 个十进制数，其范围为 0 - 255， 用(".")分割。
比如，172.16.254.1；

同时，IPv4 地址内的数不会以 0 开头。比如，地址 172.16.254.01 是不合法的。

IPv6 地址由 8 组 16 进制的数字来表示，每组表示 16 比特。这些组数字通过 (":")分割。
比如,  2001:0db8:85a3:0000:0000:8a2e:0370:7334 是一个有效的地址。
而且，我们可以加入一些以 0 开头的数字，字母可以使用大写，也可以是小写。
所以， 2001:db8:85a3:0:0:8A2E:0370:7334 也是一个有效的 IPv6 address地址
(即，忽略 0 开头，忽略大小写)。

然而，我们不能因为某个组的值为 0，而使用一个空的组，以至于出现 (::) 的情况。
比如， 2001:0db8:85a3::8A2E:0370:7334 是无效的 IPv6 地址。

同时，在 IPv6 地址中，多余的 0 也是不被允许的。
比如， 02001:0db8:85a3:0000:0000:8a2e:0370:7334 是无效的。
*/
func validIPAddress(IP string) string {
	if validIPv4Address(IP) {
		return "IPv4"
	}
	if validIPv6Address(IP) {
		return "IPv6"
	}
	return "Neither"
}

func validIPv4Address(IP string) bool {
	strArr := strings.Split(IP, ".")
	if len(strArr) != 4 {
		return false
	}
	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil || num > 255 || num < 0 { //注意：err != nil
			return false
		}
		newStr := fmt.Sprint(num)
		if str != newStr {
			return false
		}
	}
	return true
}

func validIPv6Address(IP string) bool {
	strArr := strings.Split(IP, ":")
	if len(strArr) != 8 {
		return false
	}
	for _, str := range strArr {
		if len(str) == 0 || len(str) > 4 {
			return false
		}
		for i := 0; i < len(str); i++ {
			if !(str[i] >= '0' && str[i] <= '9') &&
				!(str[i] >= 'a' && str[i] <= 'f') &&
				!(str[i] >= 'A' && str[i] <= 'F') {
				return false
			}
		}
	}
	return true
}

// func validIPAddress(IP string) string {
// 	if validIPv4Address(IP) {
// 		return "IPv4"
// 	}
// 	if validIPv6Address(IP) {
// 		return "IPv6"
// 	}
// 	return "Neither"
// }

// func validIPv4Address(IP string) bool {
// 	strArr := strings.Split(IP, ".")
// 	if len(strArr) != 4 {
// 		return false
// 	}
// 	for _, str := range strArr {
// 		if num, err := strconv.Atoi(str); err != nil || num > 255 || num < 0 {
// 			return false
// 		} else if strconv.Itoa(num) != str {
// 			return false
// 		}
// 	}
// 	return true
// }
// func validIPv6Address(IP string) bool {
// 	strArr := strings.Split(IP, ":")
// 	if len(strArr) != 8 {
// 		return false
// 	}
// 	re := regexp.MustCompile(`^([0-9]|[a-f]|[A-F])+$`)
// 	for _, str := range strArr {
// 		if len(str) == 0 || len(str) > 4 {
// 			return false
// 		}
// 		if !re.MatchString(str) {
// 			return false
// 		}
// 	}
// 	return true
// }

// 作者：xilepeng
// 链接：https://leetcode-cn.com/problems/validate-ip-address/solution/go-jian-dan-jie-fa-zheng-ze-jie-fa-by-xi-pdgj/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。