package main

type LRUCache struct {
	cache map[int]*LRUNode
	cap   int
	head  *LRUNode
	end   *LRUNode
}

type LRUNode struct {
	key   string
	value string
	pre   *LRUNode
	end   *LRUNode
}

func Constract(cap int) *LRUCache {
	head, end := new(LRUNode), new(LRUNode)
	head.next = end
	end.pre = head
	c := make(map[string]*LRUNode)
	return &LRUCache{
		cache: c,
		cap:   cap,
		head:  head,
		end:   end,
	}
}

func (this *LRUCache) Get(key string) string {
	c, ok := this.cache[key]
	if ok {
		c.remove()
		c.moveHead(this.head)
		return c.value
	}
	return "-1"
}

func (this *LRUCache) Put(key, value string) {
	c, ok := this.cache[key]
	if ok {
		this.cache[key] = n
                n.remove()
                n.moveHead(this.head)
	}
	n := newLRUNode(key, value)
	if len(this.cache) < this.cap {
		this.cache[key] = n
		n.remove()
		n.moveHead(this.head)
	} else {
		delete(this.cache, this.end.key)
		this.end.remove()
		this.cache[key] = n
		n.remove()
                n.moveHead(this.head)
	}
	
}

func newLRUNode(key, value string) *LRUNode {
	return &LRUNode{
		key: key,
		value: value,
	}
}

func (this *LRUNode) remove() {
	this.pre.next = this.next
	this.next.pre = this.pre
	this = this.next
}

func (this *LRUNode) moveHead(head *LRUNode) {
	head.pre = this
	this.next = head
}

func main() {
	lru := &LRUCache{
		cap: "1"
	}
	n1 := &LRUNode{
		key: "k1",
		value: "v1",
	}
        n2 := &LRUNode{
                key: "k2",
                value: "v2",
        }
        n3 := &LRUNode{
                key: "k3",
                value: "v3",
        }
	lru.Put(n1)
	lru.Put(n2)
	lru.Put(n3)
}
