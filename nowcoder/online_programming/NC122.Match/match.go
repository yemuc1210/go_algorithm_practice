package NC122


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 正则表达式匹配
 * 模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（包含0次）。
 * @param str string字符串 
 * @param pattern string字符串 
 * @return bool布尔型
*/
func match( str string ,  pattern string ) bool {
    // write code here
	n := len(str)
	m := len(pattern)
	f := make([][]bool,n+1)
	for i:=range f{
		f[i] = make([]bool, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			//分成空正则和非空正则两种
			if (j == 0) {
				f[i][j] = i == 0;
			} else {
				//非空正则分为两种情况 * 和 非*
				if (pattern[j - 1] != '*') {
					if (i > 0 && (str[i - 1] == pattern[j - 1] || pattern[j - 1] == '.')) {
						f[i][j] = f[i - 1][j - 1];
					}
				} else {
					//碰到 * 了，分为看和不看两种情况
					//不看
					if (j >= 2) {
						f[i][j] = f[i][j] || f[i][j - 2];
					}
					//看
					if (i >= 1 && j >= 2 && (str[i - 1] == pattern[j - 2] || pattern[j - 2] == '.')) {
						f[i][j] = f[i][j] || f[i - 1][j];
					}
				}
			}
		}
	}
	return f[n][m];
}