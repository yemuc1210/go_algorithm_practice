package lt677

/*
剑指 Offer II 066. 单词之和
实现一个 MapSum 类，支持两个方法，insert 和 sum：
	MapSum() 初始化 MapSum 对象
	void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。
	如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
	int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。
*/
type MapSum struct {
    d map[string]int
    p map[string]map[string]bool
}

func Constructor() MapSum {
    return MapSum{
        d: map[string]int{},
        p: map[string]map[string]bool{},
    }
}

func (this *MapSum) Insert(key string, val int)  {
    this.d[key] = val
    for i, n := 1, len(key); i <= n; i++ {
        if _, ok := this.p[key[: i]]; !ok {
            this.p[key[: i]] = map[string]bool{}
        }
        this.p[key[: i]][key] = false
    }
}

func (this *MapSum) Sum(prefix string) int {
    sum := 0
    for key := range this.p[prefix] {
        sum += this.d[key]
    }
    return sum
}


