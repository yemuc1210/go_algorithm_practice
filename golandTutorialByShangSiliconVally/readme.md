# 尚硅谷 Goland 入门到实战教程 Bilibili
redis 手册  redisdoc.com
http://doc.redisfans.com/

## 一 基本操作
 数据库基本操作

1、添加key-val  set

2、查看当前redis所有key:   keys *

3、获取key对应值: get key

4、切换redis 数据库[select index]

5、如何查看当前数据库的key-val 数量[dbsize]

6、清空当前数据库的key-val 和清空所有数据库的key-val [flushdb flushall]
## 二 Redis的Crub操操作
### 一 五大数据类型
Redis 的五大数据类型是: String(字符串) 、Hash (哈希)、List(列表)、Set(集合)
和zset(sorted set：有序集合)
#### 1、字符串
string 是redis 最基本的类型，一个key 对应一个value。
string 类型是二进制安全的。除普通的字符串外，也可以存放图片等数据。

redis 中字符串value 最大是512M

##### 修改操作
    set 如果存在就是修改
    setex key time val
    设置键值的生存时间，单位s
    mset 同时设置多个键值对
    [mset k1 v1 k2 v2]

#####  删除
    del key
##### 查询
    get key
    mget k1 k2


#### 2、`Hash` (类似map)
##### 基本介绍
Redis hash 是一个键值对集合。`var user1 map[string]string
Redis hash `是一个`string` 类型的`field` 和`value` 的映射表，`hash` 特别适合用于存储对
象。

##### 操作命令：
    （1）hset user1 name "xxx'
    hset user1 age 30
    hset user1 job "teacjse"

    可以一行
    hset user1 age 22 job "kkk"
     hset user2 name "john" age 33 job "coder"
    (2) hget user1 name 
    必须加上field名称才能查询
    (3)其他 
    hset/hget/hgetall/hdel
    (4)统计一个hash有多少元素
    hlen
    127.0.0.1:6379> hlen use2(integer) 3

    (5)查看哈希表key 中，给定域field 是否存在   hexists key field
    返回>0则存在


#### 3、`List` (列表)
列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列
表的头部（左边）或者尾部（右边）。
List 本质是个链表, List 的元素是有序的，元素的值可以重复.

看出管道
##### 基本操作
    lpush/rpush/lrange/lpop/rpop/del/
    (1) lpush key value[value ....]
    127.0.0.1:6379> lpush city nanjing shanghai beijing
    (integer) 3

    (2)lrange 【start,end】
    127.0.0.1:6379> lrange city 0 -1
    1) "beijing"
    2) "shanghai"
    3) "nanjing"

    (3)llen key
    127.0.0.1:6379> llen city
    (integer) 3

    (4)lpop  左边弹出第一个
    127.0.0.1:6379> lpop city
    "beijing"

    127.0.0.1:6379> lrange city 0 -1
    1) "shanghai"
    2) "nanjing"

    (5)rpop  右边弹出
    127.0.0.1:6379> lrange city 0 -1
    1) "shanghai"
    2) "nanjing"

    127.0.0.1:6379> rpop city
    "nanjing"

#### 4、 `set`集合介绍
Redis 的Set 是string 类型的`无序`集合。
 底层是HashTable 数据结构, Set 也是存放很多字符串元素，字符串元素是无序
的，而且元素的值`不能重复`
##### 操作
    (1)增加 sadd
    127.0.0.1:6379> sadd emails 111@souhu.com 100@qq.com  12212@gmail.com
    (integer) 3

    (2)smembers emails
    127.0.0.1:6379> smembers emails
    1) "111@souhu.com"
    2) "12212@gmail.com"
    3) "100@qq.com"

    再次新增  取出时不删除元素
    127.0.0.1:6379> sadd emails sjsh@xsjcn.com
    (integer) 1
    127.0.0.1:6379> smembers emails
    1) "sjsh@xsjcn.com"
    2) "111@souhu.com"
    3) "12212@gmail.com"
    4) "100@qq.com"
   
    (3) sismember[判断值是否是成员]

    (4)srem [删除指定值]



## 三 golangd操作redis

### 安装第三方开源库
1、使用第三方开源的redis 库: github.com/garyburd/redigo/redis

在使用Redis 前，先安装第三方Redis 库，在GOPATH 路径下执行安装指令:
D:\goproject>go get github.com/garyburd/redigo/redis

位置在$GOPATH/pkg/linux_amd64/github.com/garyburd/redigo

2、