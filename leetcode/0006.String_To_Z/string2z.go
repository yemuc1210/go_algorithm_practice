/*
输入字符串，给定行数从上往下，从左往右 z 形状排

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
*/

package leetcode

// import "fmt"

// type listNode struct{
// 	val byte
// 	next *listNode
// }

// //链表数组可以，不过有点麻烦
// func convert(s string, numRows int) string {
// 	arr_row := make([]listNode,numRows)   //申请空间

// 	index_i,inde_j,index_ch := 0,0,0    //行列坐标
// 	direction := 1      //从上往下为1    从下往上为-1
// 	for inde_j<numRows || index_i<numRows {
// 		if index_i < numRows-1 {
// 			a := &listNode{val:s[index_ch]}
// 			arr_row[index_i].next = a
// 			index_i += direction * 1
// 		}else if index_i == numRows-1{
// 			//最后一行的元素，下面是
// 		}
// 	}

// 	return ""
// }
// class Solution {
// 	public:
// 		string convert(string s, int numRows) {
// 			vector<string> temp(numRows);
// 			string res;
// 			if(s.empty() || numRows < 1) return res;
// 			if(numRows == 1) return s;
// 			for(int i = 0; i < s.size(); i++){
// 				int ans = i / (numRows-1);
// 				int cur = i % (numRows-1);
// 				if(ans % 2 == 0){  //结果为偶数或0
// 					temp[cur].push_back(s[i]); //按余数正序保存
// 				}
// 				if(ans % 2 != 0){  //结果为奇数
// 					temp[numRows-cur-1].push_back(s[i]); //按余数倒序保存
// 				}
// 			}
// 		   for(int i = 0; i < temp.size(); i++){
// 				   res += temp[i];
// 			}
// 			return res;
// 		}
// 	};
//直接二维数组    pass 
func convert(s string, numRows int)string{
	if numRows < 1 || len(s) == 0 {
		return ""
	}
	if len(s) < numRows || numRows==1{
		return s
	}
	arr := make([][]string,numRows)
	goDown := true
	index_arr_row := 0;

	n := len(s)
	for i := range arr {
		arr[i] = make([]string, n)
	}
	// fmt.Println(arr)
	cnt := 0
	//先放入一个
	arr[index_arr_row] = append(arr[index_arr_row], s[0:1])
	index_arr_row++
	for i:=1;i<n;i++{
		// fmt.Println(i,index_arr_row,s[i:i+1])
		arr[index_arr_row] = append(arr[index_arr_row], s[i:i+1])
		cnt ++
		if cnt%(numRows-1) == 0 {
			if goDown{
				goDown=false
			}else{
				goDown=true
			}
		}
		if goDown {
			index_arr_row ++
		} else{
			index_arr_row -- 
		}

	}
	var res string
	for _,arr1 := range arr {
		for _,val := range arr1 {
			res = res + val
		}
	}
	return res
}