package google

type LRUNode struct {
	key, val  int
	pre, next *LRUNode
}

type LRUCache struct {
	Capacity        int
	CacheMap        map[int]*LRUNode
	MostRecentHead  *LRUNode
	LeastRecentTail *LRUNode

	CurrentCapacity int
}

func Constructor(capacity int) LRUCache {
	mostRecentHead := &LRUNode{}
	leastRecentTail := &LRUNode{}
	mostRecentHead.next = leastRecentTail
	leastRecentTail.pre = mostRecentHead
	return LRUCache{
		Capacity:        capacity,
		MostRecentHead:  mostRecentHead,
		LeastRecentTail: leastRecentTail,
		CacheMap:        make(map[int]*LRUNode, capacity),
		CurrentCapacity: 0,
	}
}

func (this *LRUCache) RemoveFromList(node *LRUNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	return
}

func (this *LRUCache) InsertToList(node *LRUNode) {
	newNext := this.MostRecentHead.next
	this.MostRecentHead.next = node
	node.pre = this.MostRecentHead
	node.next = newNext
	newNext.pre = node
	return
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.CacheMap[key]
	if ok {
		this.RemoveFromList(v)
		this.InsertToList(v)
		return v.val
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.CacheMap[key]
	if ok {
		node.val = value
		this.RemoveFromList(node)
		this.InsertToList(node)
		return
	}

	newNode := &LRUNode{
		key: key,
		val: value,
	}
	this.InsertToList(newNode)
	this.CacheMap[key] = newNode
	this.CurrentCapacity = this.CurrentCapacity + 1
	if this.CurrentCapacity > this.Capacity {
		toDelete := this.LeastRecentTail.pre
		delete(this.CacheMap, toDelete.key)
		this.RemoveFromList(toDelete)
		this.CurrentCapacity = this.CurrentCapacity - 1
	}

	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
