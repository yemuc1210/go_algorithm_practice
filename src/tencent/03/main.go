package main

import (
	"fmt"
	// "math"
	"strconv"
)
type pair struct{
	s int   // 属性
	z int  // 战斗力
}
type elem struct{
	attack int
	defense int
}
func main() {
	var n int
	fmt.Scanln(&n)  // 编号1-n
	nums := make([]pair,n+1)
	var s string
	fmt.Scanln(&s)
	for i:=0;i<n;i++{
		num,_ := strconv.Atoi(s[i:i+1])
		nums[i+1].s = num  //属性
		nums[i+1].z = i+1
	}  // 下标[1,n]
	// fmt.Println(nums)
	// 从某个点分为两个阵营
	// pos   0<=pos<=n  则编号1,2,3,...pos的为一组  
	// 第一个营为进攻方,第二个为防守方
	// 第一个阵营战斗力综合为w  第二个战斗力为v
	// 希望差值最小,求|w-v|的最小值
	// 前缀和
	prefix := make([]elem,n +1) 
	// prefix[i] 表示i之前的攻击和战斗力的总和
	attackSum,defenseSum := 0,0
	for i:=1;i<=n;i++{
		// 除去当前的
		prefix[i].attack = attackSum
		prefix[i].defense = defenseSum
		if nums[i].s == 0 {  // 只会进攻
			attackSum += nums[i].z
		}else{  // 只会防御
			defenseSum += nums[i].z
		}
	}
	// fmt.Println("attackSum:",attackSum)
	// fmt.Println("defenseSum:",defenseSum)

	// fmt.Println(prefix)
	// pos点前的attack 和pos之后的defense需要接近
	// minWV := math.MaxInt32
	minWV := abs(0,defenseSum)
	if n == 1 {
		fmt.Println(minWV)
		return
	}
	for i:=1;i<=n;i++{
		// 判断当前pos是否合理
		// 第一阵营 [1,pos] 
		var w,v int
		if nums[i].s == 0 {
			// 第一阵营进攻
			w = prefix[i].attack + nums[i].z
			v = defenseSum - prefix[i].defense 
		}else{  // i 防守
			w = prefix[i].attack
			// 第二阵营 
			v = defenseSum - prefix[i].defense - nums[i].z
		}
		if abs(w,v) < minWV {
			minWV = abs(w,v)
		}
	}
	fmt.Println(minWV)
}
func abs(a,b int) int {
	if a>b {
		return a-b
	}
	return b-a
} 