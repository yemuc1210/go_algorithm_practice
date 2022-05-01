
// If you need to import additional packages or classes, please import here.
// 只能import Go标准库
// main函数是入口
// 通过标准输入流输入，通过标准输出流输出.

package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	// please define the input here.
	// For example: r := bufio.NewReader(os.Stdin) input, err := r.ReadString("\n")
	// please finish the function body here.
	// please define the output here. For example: fmt.Println(input)
	reader := bufio.NewReader(os.Stdin)
	var M,N int
	fmt.Scanln(&M)
	fmt.Scanln(&N)
	// 输入业务配置 N
	// 业务配置是
	jobs := []byte{}
	jobss,_,_ := reader.ReadLine()
	fmt.Println(jobss)

	for i:=0;i<len(jobss);i++{
	//	fmt.Scanf("%s%f",&jobs[i])
	//	fmt.Println(jobs[i])
		if jobss[i] != ' '{
			jobs = append(jobs, jobss[i])
		}
	}
	// 求解 最少的芯片
	var cpuID int = 0
	var resID int = 0
	var aIdx,bIdx int = 0,0
	m := make(map[int]int)
	m[0] = 4
	for i:=0;i<N;i++{
		if jobs[i] == 'A' {
			if m[aIdx] == 4 {
				aIdx = max(aIdx,bIdx)+1
			}
			num := m[aIdx]
			m[aIdx] = num+1
			cpuID = aIdx
			resID = num+1
		}else if jobs[i] == 'B' {
			if m[aIdx] > 0 {
				bIdx = max(aIdx,bIdx) + 1
			}
			m[bIdx] = 4
			cpuID = bIdx
			resID = 1
		}
		if aIdx>M || bIdx>M {
			break
		}
	}
	if cpuID > M {
		fmt.Println(0)
		fmt.Println(0)
	}else {
		fmt.Println(cpuID)
		fmt.Println(resID)
	}
	//
	//
	//// 四个A一个芯片  一个B一个芯片
	//var curUsedCpuId int
	//var resCpuID,resCpuResID int
	//cpus := make(map[int]int)  //cpuID-已经使用的资源数
	//for i:=0;i<N;i++{
	//	if jobs[i] == 'A'{
	//		// 尝试添加到之前已经有的
	//		id := canReUse(cpus)
	//		if id != -1 {
	//			// 说明可以重新使用
	//			cpus[id] ++  // 使用一个资源
	//		}else{
	//			// 新建一个
	//			if curUsedCpuId <  M {
	//				cpus[curUsedCpuId] ++  // 新建一个k-v
	//				resCpuID = curUsedCpuId
	//				resCpuResID = cpus[curUsedCpuId]
	//				curUsedCpuId++
	//			}else{
	//				resCpuResID = -1
	//				resCpuID = -1
	//				break
	//			}
	//		}
	//	}else if jobs[i] == 'B' {
	//		// 遇到B基本只能新建
	//		if curUsedCpuId < M {
	//			cpus[curUsedCpuId] = 4
	//			resCpuID = curUsedCpuId
	//			resCpuResID = cpus[curUsedCpuId]
	//			curUsedCpuId++
	//		}else{
	//			resCpuResID = -1
	//			resCpuID = -1
	//			break
	//		}
	//	}
	//}
	//  输出最后一个业务对应的芯片编号和恶芯片资源
	//if resCpuID!=-1 && resCpuResID!=-1 {
	//	fmt.Printf("%d\n%d\n",resCpuID-1,resCpuResID)
	//}else{
	//	fmt.Printf("0\n0")
	//}
}
func max(a,b int) int {
	if a>b{
		return a
	}
	return b
}
func canReUse(cpus map[int]int) int {
	for k,v := range cpus {
		if v < 4 {
			// 有资源，可以继续使用
			return k  // 返回cpu编号
		}
	}
	return -1  // 否则不能重用
}
