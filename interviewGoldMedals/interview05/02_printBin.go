package interview05

import (
	// "fmt"
	"strings"
	// "strconv"
)

//转二进制打印
func printBin(num float64) string {
	sb := strings.Builder{}   //Builder函数待了解
	sb.WriteString("0.")
	for i := 0; i < 31; i++ {
		num *= 2
		if num >= 1 {
			sb.WriteString("1")
			num -= 1
		} else {
			sb.WriteString("0")
		}
		if num == 0 {
			break
		}
	}
	if num != 0 {
		return "ERROR"
	}
	return sb.String()
}

