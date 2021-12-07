package offer58

// import "strings"

/*
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。
例如输入字符串"I am a student. "，则输出"student. a am I"。


这种结构，用栈啊
*/
func reverseWords(s string) string {
	if len(s)==0{
		return ""
	}
	//处理前置空格
	index := 0 
	for index<len(s) && s[index]==' ' {
		index ++
	}
	words := []string{}
	var res string
	for index<len(s){
		if s[index] != ' '{
			//则取单词
			var tmp []byte
			for index<len(s) && s[index] !=' ' {
				tmp = append(tmp, s[index])
				// fmt.Println(tmp)
				index ++
			}
			words = append(words, string(tmp))
		}
		for index<len(s) && s[index] ==' ' {
			index++
		}
	}
	//逆序输出
	for i:=len(words)-1;i>=0;i--{
		if i == 0{
			res+=words[i]
		}else{
			res += words[i]+" "
		}
	}
	return res
}

// //别人题解
// func reverseWords1(s string) string {
//     strSlice := strings.Fields(s)  //好家伙，直接调用分割
//     left, right := 0, len(strSlice) - 1
//     for left < right {
//         strSlice[left], strSlice[right] = strSlice[right], strSlice[left]
//         left++
//         right--
//     }

//     return strings.Join(strSlice, " ")
// }
