package lt273

import "strings"

/*
273. 整数转换英文表示
将非负整数 num 转换为其对应的英文表示。

示例 1：

输入：num = 123
输出："One Hundred Twenty Three"
示例 2：

输入：num = 12345
输出："Twelve Thousand Three Hundred Forty Five"


三个一组 ：thousand  million  billion
之后就是基准为billion，然后是多少多少billion
形式上是递归、
（1）小于20  直接得到英文表示
（2）【20，100）  首先十进制位转**ty  然后个位递归得到
（3）【100，）首先百位，然后递归

得到每一组数字，然后加上对应的单位，拼接
*/

var (
	singles = []string{"","One","Two","Three","Four","Five","Six","Seven","Eight","Nine"}
	teens = []string{"Ten","Eleven","Twelve","Thirteen","Fourteen","Fifteen","Sixteen","Seventeen",
"Eighteen","Nineteen"}
	tens = []string{"","Ten","Twenty","Thirty","Forty","Fifty","Sixty","Seventy","Eighty","Ninety"}
	thousands = []string{"","Thousand","Million","Billion"}
)
func numberToWords(num int) string {
	if num == 0{
		return "Zero"
	}
	stringBuilder := &strings.Builder{}
	var recursion func(int)
	recursion = func(num int){
		if num>=0 && num < 10{
			stringBuilder.WriteString(singles[num])  //直接把num当作下标找到对应英文
			stringBuilder.WriteByte(' ')// 空格
		} else if num >= 10 && num < 20 {
			stringBuilder.WriteString(teens[num-10])
			stringBuilder.WriteByte(' ')
		} else if num >= 20 && num < 100 {
			stringBuilder.WriteString(tens[num/10])   //十位
			stringBuilder.WriteByte(' ')
			//递归
			recursion(num % 10)
		} else if num >= 100 {
			//先百位
			stringBuilder.WriteString(singles[num/100])
			stringBuilder.WriteString(" Hundred")
			recursion(num % 100)
		}
	}
	//拼接，三个为一组，附加单位
	for i,unit := 3,int(1e9); i>=0;i--{
		if curNum := num/unit; curNum > 0 {
			num -= curNum * unit
			recursion(curNum)
			stringBuilder.WriteString(thousands[i])
			stringBuilder.WriteByte(' ')
		}
		unit /= 1000
	}
	return strings.TrimSpace(stringBuilder.String())
}