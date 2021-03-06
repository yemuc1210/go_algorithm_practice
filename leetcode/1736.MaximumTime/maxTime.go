package lt1736

/*
简单题
给你一个字符串 time ，格式为 hh:mm（小时：分钟），
其中某几位数字被隐藏（用 ? 表示）。

有效的时间为 00:00 到 23:59 之间的所有时间，
包括 00:00 和 23:59 。

替换 time 中隐藏的数字，返回你可以得到的最晚有效时间。
*/

/*
第一位：若第二位的值已经确定，且值落在区间 [4,9][4,9] 中时，第一位的值最大只能为 11，否则最大可以为 22；
第二位：若第一位的值已经确定，且值为 22 时，第二位的值最大为 33，否则为 99；
第三位：第三位的值的选取与其它位无关，最大为 55；
第四位：第四位的值的选取与其它位无关，最大为 99。

*/
func maximumTime(time string) string {
	t := []byte(time)
    if t[0] == '?' {
        if '4' <= t[1] && t[1] <= '9' {
            t[0] = '1'
        } else {
            t[0] = '2'
        }
    }
    if t[1] == '?' {
        if t[0] == '2' {
            t[1] = '3'
        } else {
            t[1] = '9'
        }
    }
    if t[3] == '?' {
        t[3] = '5'
    }
    if t[4] == '?' {
        t[4] = '9'
    }
    return string(t)


}
