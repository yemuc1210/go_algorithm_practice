package lt1185

func dayOfWeek(day,month,year int) string {
	//输入年月日 返回这是第几天
	weeks := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if month < 3 {
		month += 12
		year -= 1   //算作上一年的13 14月
	}
	c := year/100  
	y := year%100
	d := c/4  - 2*c + y + y/4 + 13*(month+1)/5 + day - 1 + 210  //210凑得30*7防止负数
	return weeks[d%7]
}

/*
蔡勒共识
c:世纪数减一  年份前两位
y：年份后两位
m 月份 取值3-14  某年的1、2月看作上一年的13 14 月
d 日数
W=D%7  一周中第几天
最后计算D可能为负值，处理：
*/



//方法2 从1971.1.1累加
func dayOfWeek1(day,month,year int) string{
	//闰年 平年判断   year % 400 == 0  || year%4==0 && year%100!=0   
	date := map[int]string{3:"Sunday", 4:"Monday", 5:"Tuesday", 6:"Wednesday", 0:"Thursday", 1:"Friday", 2:"Saturday"}
	mday := []int{31,28,31,30,31,30,31,31,30,31,30,31}
	d := 0 
	for i:=1971;i<year;i++{
		d += 365
		if i%400==0 || (i%4==0 && i%100!=0) {
			d += 1
		}
	}
	for i:= 0;i < month-1 ;i ++{
		d += mday[i]
	}
	//当该年该月了
	if month > 3 && (year%400==0 || (year%4==0 && year%100!=0)) {
		//闰年
		d += 1  //二月多一天
	}
	//else
	d += day
	return date[d%7]
}
