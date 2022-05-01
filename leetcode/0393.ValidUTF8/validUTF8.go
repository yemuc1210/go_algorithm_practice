package lt393

/*
给定一个表示数据的整数数组 data ，返回它是否为有效的 UTF-8 编码。

UTF-8 中的一个字符可能的长度为 1 到 4 字节，遵循以下的规则：
	对于 1 字节 的字符，字节的第一位设为 0 ，后面 7 位为这个符号的 unicode 码。
	对于 n 字节 的字符 (n > 1)，第一个字节的前 n 位都设为1，第 n+1 位设为 0 ，
		后面字节的前两位一律设为 10 。
		剩下的没有提及的二进制位，全部为这个符号的 unicode 码。
*/
func validUtf8(data []int) bool {
	var n int = 0
	for i := 0; i < len(data); i++ {
		if n > 0 {
			if data[i]>>6 != 2 {
				return false
			}
			n--
		} else if data[i]>>7 == 0 {
			n = 0
		} else if data[i]>>5 == 0b110 {
			n = 1
		} else if data[i]>>4 == 0b1110 {
			n = 2
		} else if data[i]>>3 == 0b11110 {
			n = 3
		} else {
			return false
		}
	}
	return n == 0
}
