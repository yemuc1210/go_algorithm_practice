package leetcode
/*给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB 
    ...

简单来说，就是一个26进制的问题
输入: 28
输出: "AB"
输入: 701
输出: "ZY"
*/

func convertToTitle(columnNumber int) string {
	//先得出各位数字
	// res := ""
	res := []byte{}

	for columnNumber>0{
		res = append(res, 'A'+byte((columnNumber-1)%26))
		columnNumber = (columnNumber-1)/26   //二十六进制   10进制是%9
	}
	for i,j:=0,len(res)-1;i<j;i,j=i+1,j-1{
		res[i],res[j]=res[j],res[i]     //逆序
	}
	return string(res)

}