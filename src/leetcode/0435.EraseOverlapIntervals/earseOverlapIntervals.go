package lt435

import "sort"

// Intervals define
type Intervals [][]int

func (a Intervals) Len() int {
	return len(a)
}
func (a Intervals) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Intervals) Less(i, j int) bool {
	for k := 0; k < len(a[i]); k++ {
		if a[i][k] < a[j][k] {
			return true
		} else if a[i][k] == a[j][k] {
			continue
		} else {
			return false
		}
	}
	return true
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Sort(Intervals(intervals))
	pre, res := 0, 1

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[pre][1] {
			res++
			pre = i
		} else if intervals[i][1] < intervals[pre][1] {
			pre = i
		}
	}
	return len(intervals) - res
}