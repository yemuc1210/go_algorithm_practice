package main


func countTexts(pressedKeys string) int {
	m := make(map[int]map[byte]int)  // 2- {a,b,c}
	m[2] = map[byte]int{'a':1,'b':2,'c':3}
	m[3] = map[byte]int{'d':1,'e':2,'f':3}
	m[4] = map[byte]int{'g':1,'h':2,'i':3}
	m[5] = map[byte]int{'j':1,'k':2,'l':3}
	m[6] = map[byte]int{'m':1,'n':2,'o':3}
	m[7] = map[byte]int{'p':1,'q':2,'r':3,'s':4}
	m[8] = map[byte]int{'t':1,'u':2,'v':3}
	m[9] = map[byte]int{'w':1,'x':2,'y':3,'z':4}

	// 解析按键信息
	// 222 可能是2 22  或 22 2  或 222
	// 或者当作排列组合题目来做
	// 每个数字得按键组合是固定得
	// dp 问题
	return 0

}