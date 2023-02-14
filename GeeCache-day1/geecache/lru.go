package geecache

import "container/list"

//LRU算法  最近最少使用的数据进行淘汰
//最近使用过的数据放在队列的头部，最近没有使用过的数据放在队列的尾部.缓存满了的时候，淘汰队列尾部的数据

type Cache struct {
	//最大存储空间
	maxbytes int64
	// 当前已使用的字节数
	nbytes int64
	//双向链表
	ll *list.List
	//key是字符串，value是双向链表中的节点
	cache map[string]*list.Element
	//  用于记录某条记录被移除时的回调函数，可以为 nil。
	OnEvicted func(key string, value Value)
}

//双向链表节点
type entry struct {
	key   string
	value Value
}

// 用于记录每条缓存占用的内存大小
type Value interface {
	Len() int
}

//用于实例化一个Cache
func New(maxbytes int64, OnEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxbytes:  maxbytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: OnEvicted,
	}
}

//查找功能
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok { //查找key是否存在
		c.ll.MoveToFront(ele)    //将节点移动到队列头部
		kv := ele.Value.(*entry) //获取节点的值
		return kv.value, true
	}
	return
}

//删除功能
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back() //获取队列尾部的节点
	if ele != nil {
		c.ll.Remove(ele)                                       //删除节点
		kv := ele.Value.(*entry)                               //获取节点的值
		delete(c.cache, kv.key)                                //删除map中的key
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len()) //更新已使用的字节数
		if c.OnEvicted != nil {                                //如果有回调函数，就调用回调函数
			c.OnEvicted(kv.key, kv.value) //回调函数
		}
	}
}

//添加修改
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok { //如果key存在
		c.ll.MoveToFront(ele)                                  //将节点移动到队列头部
		kv := ele.Value.(*entry)                               //获取节点的值
		c.nbytes += int64(value.Len()) - int64(kv.value.Len()) //更新已使用的字节数
		kv.value = value                                       //更新节点的值
	} else {
		ele := c.ll.PushFront(&entry{key, value})        //将节点添加到队列头部
		c.cache[key] = ele                               //添加到map中
		c.nbytes += int64(len(key)) + int64(value.Len()) //	更新已使用的字节数
	}
	for c.maxbytes != 0 && c.maxbytes < c.nbytes { //如果已使用的字节数大于最大存储空间
		c.RemoveOldest() //删除队列尾部的节点
	}
}

//返回已使用的字节数
func (c *Cache) Len() int {
	return c.ll.Len()
}
