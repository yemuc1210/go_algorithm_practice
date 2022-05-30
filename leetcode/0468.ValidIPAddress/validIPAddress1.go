package lt468

import "regexp"


func validIPAddress1(queryIP string) string {
	if isValidIPv4Address(queryIP) {
		return "IPv4"
	}
	if isValidIPv6Address(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func isValidIPv4Address(IP string) bool {
	// 使用正则表达式匹配
	re:= regexp.MustCompile("^(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)(\\.(1\\d{2}|2[0-4]\\d|25[0-5]|[1-9]\\d|\\d)){3}")
	return re.MatchString(IP)
}
func isValidIPv6Address(IP string) bool {
	re := regexp.MustCompile("^[0-9 a-f A-F]{1,4}(:[0-9 a-f A-F]{1,4}){7}")
	return re.MatchString(IP)
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
