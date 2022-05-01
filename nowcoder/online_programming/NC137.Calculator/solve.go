package NC137


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 返回表达式的值  整数计算器 支持加减乘+括号
 * @param s string字符串 待计算的表达式
 * @return int整型
*/
func solve( s string ) int {
    // 思路：转为后缀表达式求值
	//首先让s左右都是括号，省去讨论
	s = "("+s+")"
	ch,d := []byte{},[]int{}  //一个存运算符，一个存数字
	i :=0
	for i<len(s){
		//如果当前是运算符
		if (s[i]=='+'||s[i]=='-'||s[i]=='*') {
			//在一对括号中间运算
			if s[i]!=')'{
				ch = append(ch, s[i])
			}else{
				//cur 记录括号内结果
				cur :=0
				for ch[len(ch)-1] !='(' {
					//如果碰到乘法，不需要更新cur,只需要更新两个占
					if ch[len(ch)-1] == '*' {
						x1,x2 := d[len(d)-1],d[len(d)-2]
						d = d[:len(d)-2]
						d = append(d, x1*x2)
					}else {
						//加减法，
						sign := ch[len(ch)-1]
						b := d[len(ch)-1]
						d = d[:len(d)-1]
						if sign=='-'{
							b = -b
						}
						cur += b
					}
					ch = ch[:len(ch)-1]
				}
				//
				// 之前的过程中是每碰到一个加减法，就去数字栈中末尾元素加入cur，
                // 如果是(-2+3)就刚好处理完，如果是(2+3)，那么数字占中还剩余一个2，把它也加入cur中
				if len(d)>0 {
					cur += d[len(d)-1]
					d = d[:len(d)-1]
				}
				// 把这对括号中的运算式替换成cur，并删掉栈顶的‘(’
				d = append(d, cur)
				ch = ch[:len(ch)-1]
			}
			i++
		}else{//当前字符是数字
			cur := 0
			for s[i]>='0' && s[i]<='9'{
				cur = cur*10 + int(s[i]-'0')
				i++
			}
			d = append(d, cur)
		}
	}
	return d[len(d)-1]
}


// public int solve (String s) {
        
// 	// 为了方便计算，让所有的s都以左括号开始，右括号结束
// 	s = "(" + s + ")";
	
// 	// c用来存运算符，d用来存数字
// 	Stack<Character> c = new Stack<>();
// 	Stack<Integer> d = new Stack<>();
	
// 	int i = 0;
// 	while (i < s.length()) {
		
// 		// 当前字符为运算符
// 		if (!Character.isDigit(s.charAt(i))) {
			
// 			// 在一对括号中间进行运算
// 			if (s.charAt(i) != ')') {
// 				c.add(s.charAt(i));
// 			} else {
				
// 				// cur记录括号内的计算结果
// 				int cur = 0;
// 				while (c.peek() != '(') {
					
// 					// 如果碰到乘法，不需要更新cur，只需要更新两个栈
// 					if (c.peek() == '*') {
// 						d.add(d.pop() * d.pop());
						
// 					// 如果碰到加减法，就把两个栈尾的元素组合一下，加入cur
// 					} else {
// 						Character sign = c.peek();
// 						int b = d.pop();
// 						if (sign == '-') {
// 							b = -b;
// 						}
// 						cur += b;
// 					}
// 					c.pop();
// 				}
				
// 				// 之前的过程中是每碰到一个加减法，就去数字栈中末尾元素加入cur，
// 				// 如果是(-2+3)就刚好处理完，如果是(2+3)，那么数字占中还剩余一个2，把它也加入cur中
// 				if (!d.isEmpty()) {
// 					cur += d.pop();
// 				}
				
// 				// 把这对括号中的运算式替换成cur，并删掉栈顶的‘(’
// 				d.add(cur);
// 				c.pop();
// 			}
// 			i ++;
			
// 		// 当前字符为数字,就把数字添加到数字栈中
// 		} else {
// 			int cur = 0;
// 			while (Character.isDigit(s.charAt(i))) {
// 				cur = 10 * cur + Integer.valueOf(s.charAt(i) - 48);
// 				i ++;
// 			}
// 			d.add(cur);
// 		}
// 	}
// 	return d.pop();
// }