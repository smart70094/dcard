package ad

import "sync"

var adCache = NewCache(1000)

type Cache struct {
	mutex sync.RWMutex
	cache map[string]interface{}
	keys  []string // 用于维护项目的顺序
	limit int      // 缓存的最大容量
}

func NewCache(limit int) *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
		keys:  make([]string, 0),
		limit: limit,
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// 如果超过了最大容量，删除最旧的项目
	if len(c.keys) >= c.limit {
		oldestKey := c.keys[0]
		delete(c.cache, oldestKey)
		c.keys = c.keys[1:]
	}

	// 添加新项目到缓存中
	c.cache[key] = value
	c.keys = append(c.keys, key)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, ok := c.cache[key]
	return value, ok
}
