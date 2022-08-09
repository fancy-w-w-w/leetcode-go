package specialquestion

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// 带过期时间的lru链表
type value struct {
	data        []byte
	lruPos      *list.Element
	expiredTime int64
}

type lruExpiredCache struct {
	maxSize int
	data    map[interface{}]*value
	lck     *sync.Mutex
	lru     *list.List
}

func NewCache(maxLength int) *lruExpiredCache {
	if maxLength <= 0 {
		panic("maxLength can not < 0 in NewCache")
	}
	return &lruExpiredCache{
		maxSize: maxLength,
		data:    make(map[interface{}]*value),
		lck:     new(sync.Mutex),
		lru:     list.New(),
	}
}

func (c *lruExpiredCache) Set(key interface{}, data []byte) {
	c.lck.Lock()
	defer c.lck.Unlock()
	//c.updateKey(key) //这里不用更新状态，因为数据即将被覆盖
	if val, found := c.data[key]; found {
		c.deleteLruItem(val.lruPos)       //删除原先在list中的位置
		val.lruPos = c.updateNewItem(key) // 追加到list末尾，更新位置标示
		val.expiredTime = -1
		c.data[key] = val
	} else {
		var pos *list.Element
		if len(c.data) < c.maxSize {
			pos = c.updateNewItem(key) //直接追加到list末尾
		} else {
			c.deleteLruItem(c.lru.Front()) //删除最久未使用的
			pos = c.updateNewItem(key)     //将本次key更新到list的末尾
		}
		val := &value{
			data:        data,
			lruPos:      pos,
			expiredTime: -1,
		}
		c.data[key] = val
	}
}

//lazy删除，等到实际用的时候判断key是否应该被删除。
func (c *lruExpiredCache) updateKey(key interface{}) {
	if val, found := c.data[key]; found {
		if val.expiredTime <= time.Now().Unix() {
			c.deleteLruItem(val.lruPos) //在map中和list中均需要删除
		}
	}
}

func (c *lruExpiredCache) Exist(key interface{}) bool {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.updateKey(key)
	_, found := c.data[key]
	return found
}

func (c *lruExpiredCache) SetWithExpiredTime(key interface{}, data []byte, expiredSeconds int64) {
	if expiredSeconds <= 0 {
		return
	}
	c.lck.Lock()
	defer c.lck.Unlock()
	c.updateKey(key)
	expiredTime := time.Now().Unix() + expiredSeconds
	if val, found := c.data[key]; found {
		c.deleteLruItem(val.lruPos)       //删除原先在list中的位置
		val.lruPos = c.updateNewItem(key) // 追加到list末尾，更新位置标示
		val.expiredTime = expiredTime
		c.data[key] = val
	} else {
		var pos *list.Element
		if len(c.data) < c.maxSize {
			pos = c.updateNewItem(key) //直接追加到list末尾
		} else {
			c.deleteLruItem(c.lru.Front()) //删除最久未使用的
			pos = c.updateNewItem(key)     //将本次key更新到list的末尾
		}
		val := &value{
			data:        data,
			lruPos:      pos,
			expiredTime: expiredTime,
		}
		c.data[key] = val
	}
}

func (c *lruExpiredCache) Get(key interface{}) (exist bool, data []byte) {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.updateKey(key)
	if val, found := c.data[key]; found {
		data = val.data
		c.deleteLruItem(val.lruPos)       //删除原先在list中的位置
		val.lruPos = c.updateNewItem(key) // 追加到list末尾，更新位置标示
		c.data[key] = val
		exist = true
	}
	return
}

func (c *lruExpiredCache) deleteLruItem(pos *list.Element) (item *list.Element) {
	c.lru.Remove(pos)
	delete(c.data, pos.Value)
	return
}

func (c *lruExpiredCache) updateNewItem(key interface{}) (item *list.Element) {
	item = c.lru.PushBack(key)
	return
}

// lazy更新的话，这样Length就会不准了
//func (c *lruExpiredCache) Length() int {
//	c.lck.Lock()
//	defer c.lck.Unlock()
//	return len(c.data)
//}

func (c *lruExpiredCache) DebugShowMapData() {
	c.lck.Lock()
	defer c.lck.Unlock()
	fmt.Println("=== map ===")
	for k, v := range c.data {
		fmt.Println(k, v)
	}
	fmt.Println("=== map over ===")
}
func (c *lruExpiredCache) DebugShowLruList() {
	c.lck.Lock()
	defer c.lck.Unlock()
	fmt.Println("=== list ===")
	for v := c.lru.Front(); v != nil; v = v.Next() {
		fmt.Print(v.Value, " ")
	}
	fmt.Println("\n=== list over ===")
}

func (c *lruExpiredCache) Delete(key interface{}) {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.updateKey(key)
	if val, found := c.data[key]; found {
		c.deleteLruItem(val.lruPos)
	}
}
