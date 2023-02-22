package geecache

import (
	"fmt"
	"geecache/lru"
	"sync"
)

//可以看为缓存的命名空间，将不同类型的缓存分组存储
type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mu     sync.RWMutex //读锁
	groups = make(map[string]*Group)
)

//
type Getter interface {
	Get(key string) ([]byte, error)
}

//	 定义一个回调函数
type GetterFunc func(key string) ([]byte, error)

// 该函数的作用是将回调函数转换为接口类型
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

//添加数据
func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}

	c.lru.Add(key, value)

}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}

// 创建Group实例
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		fmt.Println("nil getter!")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

//获取缓存分组命名空间
func GetGroup(name string) *Group {
	mu.Lock()
	g := groups[name]
	defer mu.Unlock()
	return g
}
