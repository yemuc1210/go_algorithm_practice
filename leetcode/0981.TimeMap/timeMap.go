package leetcode

import "sort"

/*
创建一个基于时间的键值存储类 TimeMap，它支持下面两个操作：

1. set(string key, string value, int timestamp)

    存储键 key、值 value，以及给定的时间戳 timestamp。

2. get(string key, int timestamp)

    返回先前调用 set(key, value, timestamp_prev) 所存储的值，其中 timestamp_prev <= timestamp。
    如果有多个这样的值，则返回对应最大的  timestamp_prev 的那个值。
    如果没有值，则返回空字符串（""）。

*/
type data struct{
	timestamp int
	value string
}
type TimeMap struct {
	m map[string][]data
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{map[string][]data{}}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	this.m[key] = append(this.m[key], data{timestamp,value})
}

//查找时间戳小于等于给定的，如果有多个，返回最大的，没有则返回空
func (this *TimeMap) Get(key string, timestamp int) string {
	data := this.m[key]
	//Search uses binary search to find and return the smallest index i in [0, n) 
	//at which f(i) is true
	i:=sort.Search(len(data),func(i int)bool{
		return timestamp < data[i].timestamp
	})
	if i > 0 {
        return data[i-1].value
    }
    return ""
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */