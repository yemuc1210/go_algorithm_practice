package meituan001

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
小美是美团的前端工程师，为了防止系统被恶意攻击，小美必须要在用户输入用户名之前做一个合法性检查，一个合法的用户名必须满足以下几个要求：

用户名的首字符必须是大写或者小写字母。
用户名只能包含大小写字母，数字。
用户名需要包含至少一个字母和一个数字。
如果用户名合法，请输出 "Accept"，反之输出 "Wrong"。


格式：


输入：
- 输入第一行包含一个正整数 T，表示需要检验的用户名数量。
- 接下来有 T 行，每行一个字符串 s，表示输入的用户名。
输出：
- 对于每一个输入的用户名 s，请输出一行，即按题目要求输出一个字符串。
示例：


输入：
     5
     Ooook
     Hhhh666
     ABCD
     Meituan
     6666
输出：
     Wrong
     Accept
     Wrong
     Wrong
     Wrong

*/
func isValidName(T int, names []string) []string {
	res := []string{}
	for i:=0;i<T;i++{
		if isValid(names[i]){
			res = append(res, "Accept")
		}else{
			res = append(res, "Wrong")
		}
	}
	return res
}

//也可以使用正则包 regexp
func isValid(name string)bool{
	//返回用户名是否合法
	if len(name)<2{
		return false
	}
	isDigit := false
	isUpper,isLower := false,false
	for i:=0;i<len(name);i++{
		if i==0 {
			//首字符必须为字母
			if (name[i]>'a' && name[i]<'z' )  {
				isLower = true
				continue
			}else if (name[i]>'A'&&name[i]<'Z'){
				isUpper = true
				continue
			}else{
				return false
			}
		}
		//首字符是大小写字母，那么剩下的必须出现数字     那么大小写都需要出现？
		if name[i]>'0' && name[i]<'9'{
			isDigit = true
		}else{   //不是数字，判断是否是字母
			if (name[i]>'a' && name[i]<'z' ) {	
				isLower = true
			}else if (name[i]>'A'&&name[i]<'Z') {
				isUpper = true
			}else{ //非字符，非数字
				return false
			}
		}
	}
	return isDigit && isLower && isUpper
}

func main(){
	input := bufio.NewScanner(os.Stdin)

	var size int
	if input.Scan() {
		size,_ = strconv.Atoi(input.Text())
	}
	var names []string
	for i:=0;i<size;i++{
		if input.Scan() {
			names[i] = input.Text()
		}
	}


	res := isValidName(size,names)
	for _,val := range res{
		fmt.Println(val)
	}
}