package NC149
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * kmp算法
 * 计算模板串S在文本串T中出现了多少次
 * @param S string字符串 模板串
 * @param T string字符串 文本串
 * @return int整型
*/
// func  kmp(mainString string, subString string) int{
// 	mainIdx :=0;
//     subIdx :=0;
//     mainLen :=len(mainString)
//     subLen := len(subString)
//     next := computeNextArray(subString);
//     for  {
//         if mainIdx>=mainLen  || subIdx >= subLen{
//             //当子串匹配完, 说明主串中包含子串, 需要结束
//             //当主串匹配完, 不管子串有没有结束, 都应该不再匹配
//             break;
//         }

//         if mainString[mainIdx] == subString[subIdx]{
//             //字符匹配的时候, 主串和子串都往后走
//             mainIdx ++;
//             subIdx ++;
//         }else{
//             //不相等的时候, 子串切换到next数组中对应的前一个元素对应的, 需要位移的下标, 进行重新匹配
//             if subIdx != 0{
//                 subIdx = next[subIdx-1]
//             }else{
//                 //当主串元素与子串元素不一样, 且子串回退到了首字符时, 检查主串的下一个字符
//                 mainIdx++
//             }

//         }
//     }

//     //判断结果
//     if subIdx>=subLen{
//         //说明子串匹配完了, 子串在主串中存在
//         return mainIdx - subLen
//     }
//     //其他情况下, 说明子串没走完, 返回 -1
//     return -1

// }



// //computeNextArray 用于计算子串的 next 数组
// func computeNextArray(subString string) []int {
//     next := make([]int, len(subString))
//     index :=0; 
//     i := 1;
//     for i < len(subString){
//         if subString[i] == subString[index]{
//             next[i] = index + 1
//             i ++ ;
//             index ++;
//         }else{
//             //不相等的时候
//             if index != 0{
//                 //当index != 0 的时候, 把 next 中, index 前一个元素对应的 next中的值, 赋给 index  
//                 index = next[index -1]
//             }else{
//                 //当index = 0 的时候,  subString[i] 和 subString[0] , 也就是subString 首字母不相同, 则i向后走
//                 i++
//             }
//         }
//     }
//     return next
// }


func kmp(S,T string) int {
	if len(S)==0 || len(S) >len(T) {
		return 0  //S是子串
	}
	res,next := 0, make([]int,len(S)+1) //结果和next数组，next[x]的意义是长度为x的前子串最长相同前后缀长度
	next[0] = -1 //第0位设置成-1表示需要右移被匹配字符指针
	for pre,cur:=-1,0;cur<len(S);{
		if pre == -1 || S[pre] == S[cur] {//第一次
			pre++
			cur++
			next[cur] = pre //模板指针指向（含当前字符）最长相同前缀后一个字符
		}else{
			pre = next[pre]  //字符不相同时左指针更新为，（不含该字符之前的字符串）最长相同前缀的后一位（若该字符为S[0]则值为-1表示需要右移cur）
		}
	}
	for pattern,cur:=0,0;cur<len(T);{
		if pattern==-1 || S[pattern]==T[cur]{
			pattern++//模板串指针右移
			cur++//文本串指针右移
			if pattern==len(S){
				res ++
				pattern = next[pattern] //  //模板指针指向（含当前字符）最长相同前缀后一个字符
			}
		}else{
			pattern = next[pattern] //不相同时模板指针指向（不含该字符之前的字符串）最长相同前缀的后一位（若该字符为S[0]则值为-1表示需要右移文本串指针）
		}
	}
	return res
}