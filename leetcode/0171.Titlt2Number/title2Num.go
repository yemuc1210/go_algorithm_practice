package lt171

import "strings"

/*
根据列名称，返回列序号  26进制
*/

func titleToNumber(columnTitle string) int {
	if len(columnTitle) == 0{
		return 0
	}
	//一般来说，还需要将字母全部大写
	s := strings.ToUpper(columnTitle)
	res := 0
	for len(s) > 0 {
		lastByte := s[0]
		s = s[1:]
		res = res*26 + int(lastByte-'A'+1)
	}
	return res
}	
