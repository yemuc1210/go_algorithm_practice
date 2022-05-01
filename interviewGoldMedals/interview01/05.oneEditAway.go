package interview01
/*
面试题 01.05. 一次编辑
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 

给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。



当 len(first) == len(second) 时，如果只有一个字符不同，说明可以通过变换，实现字符串相同
当 | len(first) - len(second) | == 1 时，把长的字符串去掉一个，剩余部分和短的字符串一致，说明两字符串可通过变换，实现字符串相同
当 | len(first) - len(second) | > 1 时，两字字符串不可以通过变换，实现字符串相同


*/
func oneEditAway(first string, second string) bool {
	if len(first) > len(second) {
		first,second = second,first  //first更短
	}
	if len(second) - len(first) > 1{
		return false
	}
	for i:=0;i<len(first);i++ {
		if first[i] != second[i] {
			//消耗一个编辑
			//如果是删除
			if first[i:] == second[i+1:]{
				return true
			}
			//如果是修改
			if first[i+1:] == second[i+1:]{
				return true
			}

			//否则
			return false
		}
	}
	return true
}