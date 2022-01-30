package NC147

import "sort"

/*
* 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
* 计算成功举办活动需要多少名主持人
* @param n int整型 有n个活动
* @param startEnd int整型二维数组 startEnd[i][0]用于表示第i个活动的开始时间，startEnd[i][1]表示第i个活动的结束时间
* @return int整型
 */

/*
循环遍历
对活动开始时间进行排序
对活动结束时间进行排序
starts[start] >= ends[end]时end++;
否则count++即需要增加一个主持人
*/
func minmumNumberOfHost( n int ,  startEnd [][]int ) int {
   // write code here
   starts := make([]int,len(startEnd))
   ends := make([]int, len(startEnd))

   for i:=0;i<len(startEnd);i++{
	   starts[i] = startEnd[i][0]
	   ends[i] = startEnd[i][1]
   }
   //排序
   sort.Ints(starts)
   sort.Ints(ends)

   cnt:=0
   end := 0

   for start:=0;start<len(startEnd);start++{
	   if starts[start] >= ends[end] {
		   end++
	   }else{
		   cnt++
	   }
   }
   return cnt

}