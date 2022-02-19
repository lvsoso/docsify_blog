package main

import (
	"container/list"
)

//Cache 是一个 LRU Cache，注意它并不是并发安全的
type Cache struct {
	//MaxEntries 是 Cache 中实体的最大数量，0 表示没有限制
	MaxEntries int

	//OnEvicted 是一个可选的回调函数，当一个实体从 Cache 中被移除时执行
	OnEvicted func(key Key, value interface{})

	//ll是一个双向链表指针，执行一个 container/list 包中的双向链表
	ll *list.List

	//cache 是一个 map，存放具体的 k/v 对，value 是双向链表中的具体元素，也就是 *Element
	cache map[interface{}]*list.Element
}

//key 是接口，可以是任意类型
type Key interface{}

//一个 entry 包含一个 key 和一个 value，都是任意类型
type entry struct {
	key   Key
	value interface{}
}

//创建一个 LRU Cache。maxEntries 为 0 表示缓存没有大小限制
func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

//向 Cache 中插入一个 KV
func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

//传入一个 key，返回一个是否有该 key 以及对应 value
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

//从 Cache 中删除一个 KV
func (c *Cache) Remove(key Key) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

//从 Cache 中删除最久未被访问的数据
func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

//从 Cache 中删除一个元素，供内部调用
func (c *Cache) removeElement(e *list.Element) {
	//先从 list 中删除
	c.ll.Remove(e)

	kv := e.Value.(*entry)

	//再从 map 中删除
	delete(c.cache, kv.key)

	//如果回调函数不为空则调用
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

//获取 Cache 当前的元素个数
func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.ll.Len()
}

//清空 Cache
func (c *Cache) Clear() {
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.ll = nil
	c.cache = nil
}

func main() {
	cache := New(3)
	cache.Add("a", 1)
	println(cache.ll.Len())
	cache.Add("b", 2)
	println(cache.ll.Len())

	cache.Add("c", 3)
	println(cache.ll.Len())

	cache.Add("d", 4)
	println(cache.ll.Len())

}
