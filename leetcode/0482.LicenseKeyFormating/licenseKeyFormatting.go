package lt482

/*
482. 密钥格式化
有一个密钥字符串 S ，只包含字母，数字以及 '-'（破折号）。其中， N 个 '-' 将字符串分成了 N+1 组。

给你一个数字 K，请你重新格式化字符串，使每个分组恰好包含 K 个字符。
特别地，第一个分组包含的字符个数必须小于等于 K，但至少要包含 1 个字符。
两个分组之间需要用 '-'（破折号）隔开，并且将所有的小写字母转换为大写字母。

给定非空字符串 S 和数字 K，按照上面描述的规则进行格式化。
*/

/*
原始license每一组长度不一定
*/
func licenseKeyFormatting(s string, k int) string {
	//首先  健壮性保证  处理特殊情况
	if k==0 {
		return ""
	}
	// if len(s) <= 1 {
	// 	return s
	// }

	//获取所有字符
	chars := []byte{}
	for i:=0;i<len(s);i++ {
		if s[i] != '-' {
			//变为大写并保存
			if s[i]>='a' && s[i]<='z' {
				chars = append(chars, s[i]-32)
			}else{
				chars = append(chars, s[i])
			}
			
		}
	}

	//格式化，要求除了第一组小于等于k，其他等于k
	//计算第一组的大小
	num := len(chars)%k
	if num == 0 {
		num = k
	}

	res := []byte{}
	tmpk := k
	for i:=0;i<len(chars);i++{
		if num>0{
			//首先存第一组
			res = append(res, chars[i])
			num--
			if num==0 && i<len(chars)-1 {
				res = append(res, '-')
			}
		}else{
			//开始存其他组
			res = append(res, chars[i])
			tmpk--
			if tmpk==0 && i<len(chars)-1{  //额外考虑最后一组的情况
				res = append(res, '-')
				tmpk = k
			}
		}
	}

	return string(res)
}