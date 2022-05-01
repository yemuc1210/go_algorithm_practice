package offer44

import "strconv"

/*
数字以0123456789101112131415…的格式序列化到一个字符序列中。
在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
请写一个函数，求任意第n位对应的数字。

数位n
数字num
位数digit
每一个位数的起始start
数字范围    位数    数字数量    数位数量
1~9          1       9          9
10~99        2       90         180
100~999      3       900        2700

start~end    digit   9*start     9*start*digit


求解分为三步
1 确定n所在数字的位数    记为digit
2 确定n所在的数字     num
3 确定n是num中哪个数位
*/
func findNthDigit(n int) int {
    digit,digitNum,count := 1,1,9       //digitNum start
    for n>count{
        n -= count
        digit++
        digitNum *= 10
        count = 9*digit*digitNum
    }
    num := digitNum + (n-1)/digit

    index := (n-1)%digit

    numStr := strconv.Itoa(num)
    return int(numStr[index]-'0')
}
