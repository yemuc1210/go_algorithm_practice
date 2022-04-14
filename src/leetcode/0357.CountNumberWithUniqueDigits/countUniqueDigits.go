package lt357

func countNumbersWithUnqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	prev := 10 // 上一个状态
	for i, cur := 9, 9; n-1 > 0; {
		n--
		tmp := cur * i
		prev += tmp
		i--
	}
	return prev
}
