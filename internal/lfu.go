package internal

import "container/list"

type LFUCache struct {
	kItems map[int]*list.Element //map1 存储值对应的
	fItems map[int]*list.List
	cap    int
	minFre int
}

type entry struct {
	key, value int
	freq       int
}

func ConstructorLfu(capacity int) LFUCache {
	lfu := LFUCache{
		kItems: make(map[int]*list.Element),
		fItems: make(map[int]*list.List),
		cap:    capacity,
		minFre: 1,
	}
	lfu.fItems[lfu.minFre] = list.New() //这个频率是一定会用到的，提前申请好
	return lfu
}

func (this *LFUCache) Get(key int) int {
	if len(this.kItems) == 0 {
		return -1
	}

	node, ok := this.kItems[key]
	if !ok {
		return -1
	}
	value := node.Value.(*entry).value
	this.nodeExec(node)
	return value
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}

	node, ok := this.kItems[key]
	if ok { //该键值已经存在
		node.Value.(*entry).value = value
		this.nodeExec(node)
		return
	}

	//该键值不存在

	if len(this.kItems) >= this.cap { //如果lfu满了
		this.remove()
	}
	kv := &entry{key: key, value: value, freq: 1}
	node = this.fItems[kv.freq].PushFront(kv)
	this.kItems[key] = node
	this.minFre = 1
}

func (this *LFUCache) remove() {
	l := this.fItems[this.minFre]
	node := l.Back()
	l.Remove(node)
	delete(this.kItems, node.Value.(*entry).key)
}

func (this *LFUCache) nodeExec(node *list.Element) {
	//原频率中删除
	kv := node.Value.(*entry)
	oldList := this.fItems[kv.freq]
	oldList.Remove(node)

	//更新minfreq
	if oldList.Len() == 0 && this.minFre == kv.freq {
		this.minFre++
	}

	//放入新的频率链表
	kv.freq++
	if _, ok := this.fItems[kv.freq]; !ok {
		this.fItems[kv.freq] = list.New()
	}
	newList := this.fItems[kv.freq]
	node = newList.PushFront(kv)
	this.kItems[kv.key] = node
}
