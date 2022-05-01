package interview01

import "bytes"

/*

面试题 01.03. URL化

URL化。编写一种方法，将字符串中的空格全部替换为%20。
假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）


*/
func replaceSpaces(s string, length int) string {
	//获取单词，遇到空格就替换 append
	if len(s) == 0 {
		return ""
	}
	//处理前置空格
	index := 0
	// for index < len(s) && s[index] == ' ' && length>0 {
	// 	length -- 
	// 	index++
	// }
	// if length==0{
	// 	//说明字符串全是零
		
	// }
	//提取单词
	var res string
	
	for index < len(s) && length > 0{
		var word []byte
		if s[index] != ' ' {
			//获得非空格字符

			for index<len(s) && s[index] != ' ' && length>0{
				word = append(word, s[index])
				index ++
				length --
			}
			res += string(word)
		}else {
			//否则，遇到空格
			for index <len(s) && s[index]==' ' && length>0{
				length--
				index++
				res += "%20"
			}
			
		}
		
	}

	return res
}

func replaceSpaces1(S string, length int) string {
	runeS := []rune(S)
	i := length - 1
	j := len(S) - 1

	for i >= 0 {
		if runeS[i] == ' ' {
			runeS[j] = '0'
			j--
			runeS[j] = '2'
			j--
			runeS[j] = '%'
			j--
		} else {
			runeS[j] = runeS[i]
			j--
		}
		i--

	}
	return string(runeS[j+1 : len(S)])
}

func replaceSpaces2(S string, length int) string {
    var buf bytes.Buffer

    for i:=0; i<length; i++{
        if S[i] == ' ' {
            buf.WriteString("%20")
        }else{
            buf.WriteByte(S[i])
        }
    }

    return buf.String()
}
