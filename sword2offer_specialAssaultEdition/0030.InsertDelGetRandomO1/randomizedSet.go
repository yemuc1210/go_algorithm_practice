package offerS30

import "math/rand"

/*
剑指 Offer II 030. 插入、删除和随机访问都是 O(1) 的容器
设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构：

insert(val)：当元素 val 不存在时返回 true ，并向集合中插入该项，否则返回 false 。
remove(val)：当元素 val 存在时返回 true ，并从集合中移除该项，否则返回 false 。
getRandom：随机返回现有集合中的一项。每个元素应该有 相同的概率 被返回。
*/

type RandomizedSet struct {
	dict map[int]int
	list []int
}

func Constructor() RandomizedSet {
	dict := make(map[int]int)
	var list []int
	return RandomizedSet{dict, list}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, has := this.dict[val]; has {
		return false
	}
	this.dict[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, has := this.dict[val]; !has {
		return false
	} else {
		lastIdx := len(this.list) - 1
		lastVal := this.list[lastIdx]
		this.list[idx] = this.list[lastIdx]
		this.dict[lastVal] = idx
		this.list = this.list[:lastIdx]
		delete(this.dict, val)
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}


