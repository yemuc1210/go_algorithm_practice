package lt1154
/*
1154. 一年中的第几天
给你一个字符串 date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。请你计算并返回该日期是当年的第几天。

通常情况下，我们认为 1 月 1 日是每年的第 1 天，1 月 2 日是每年的第 2 天，依此类推。每个月的天数与现行公元纪年法（格里高利历）一致。
*/
import "time"

func dayOfYear(date string) int {
	first := date[:4] + "-01-01"  //第一天
	firstDay, _ := time.Parse("2006-01-02", first)
	dateDay, _ := time.Parse("2006-01-02", date)
	duration := dateDay.Sub(firstDay)
	return int(duration.Hours())/24 + 1
}
