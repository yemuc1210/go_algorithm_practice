package offer20


/*
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
数值（按顺序）可以分成以下几个部分：
	若干空格
	一个 小数 或者 整数
	（可选）一个 'e' 或 'E' ，后面跟着一个 整数
	若干空格

小数（按顺序）可以分成以下几个部分：
	（可选）一个符号字符（'+' 或 '-'）
	下述格式之一：
		至少一位数字，后面跟着一个点 '.'
		至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
		一个点 '.' ，后面跟着至少一位数字
整数（按顺序）可以分成以下几个部分：
	（可选）一个符号字符（'+' 或 '-'）
	至少一位数字
部分数值列举如下：
["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
部分非数值列举如下：
["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]


*/

//说明，长度在【1，20】,仅包含英文字母，数字 +- . ' '
func isNumber(s string) bool {
	if len(s)==0{
		return false
	}
	index:=0
	for index<len(s)&&s[index]==' '{
		index++
	}//去除数值前的空格干扰
	if index<len(s)&& (s[index] == '+' || s[index]=='-'){
		index++
	}  //加减号略过
	if index<len(s)&&(s[index] =='+'||s[index]=='-'){
		return false   //又出现加减号，非数值
	}else if index<len(s)&& (s[index]>='0'&&s[index]<='9'){
		//整数，第一个是数字   不过可能出现12e/1a3.14这种非数值情况
		return isDeli(s[index:]) || isInteger(s[index:])	
	}
	return index<len(s) && isDeli(s[index:])
}
//判断小数  假设已经去除空格和正负号
func isDeli(s string)bool{
	wflag := false//后续的空格
	for i:=0;i<len(s);i++{
		if !wflag &&(s[i]>='0'&&s[i]<='9'){
			continue
		}else if !wflag && s[i]=='.'{
			if i==0 && i+1==len(s){
				return false   //
			}
			//后面需要跟数字，并且不能有小数点  e
			j:=i+1
			if j<len(s)&&(s[j]=='+'||s[j]=='-'){
				return false
			}
			flag := false //用于记录周后是否是连续的空格，如果一直是空格后面又出行数字，错误
			flagD := true
			for ;j<len(s);j++{
				// fmt.Printf("s=%v, s[j]=%v\n",s,s[j])
				if !flag && flagD && (s[j]>='0'&&s[j]<='9'){
					continue
				}else if s[j]==' '&& j>=i+1{
					flag=true
					flagD=false
					continue
				}else{
					return false
				}  //e、E之后只允许出现整数
			}
			//如果没有return 说明e之后确实是整数
			return true && (flag && !flagD) //直接return 吧
		}else if  s[i]==' '&& i>0{
			wflag=true
			continue
		}else {
			return false
		}
	}
	return true

}
//假设已经去除了前置空字符和正负号
func isInteger(s string)bool{
	wflag := false
	for i:=0;i<len(s);i++{
		// fmt.Println(s[i:i+1])
		if !wflag && (s[i] >='0' && s[i]<='9'){
			continue
		}else if !wflag && (s[i]=='e'||s[i]=='E'){
			//判断后面是否是整数	
			// fmt.Printf("出现e，i+1=%v\n",i+1)
			if i+1 == len(s){//e之后没有空格
				return false
			}
			j:=i+1
			if s[j]=='+'||s[j]=='-'{
				j++
			}
			flag := false //用于记录之前是否是空格
			for ;j<len(s);j++{
				if !flag && (s[j]>='0'&&s[j]<='9'){
					continue
				}else if s[j]==' '&& j>i+1{
					flag=true
					continue
				}else{
					return false
				}  //e、E之后只允许出现整数
			}
			//如果没有return 说明e之后确实是整数
			return true    //直接return 吧
		}else if s[i]==' '&& i>0{   //整数后面解空格
			wflag=true
			continue
		}else{
			return false
		}
	}
	return true
}