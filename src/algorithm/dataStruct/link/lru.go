package main

import "fmt"

func main() {
	lru := &LRU{make(map[string]interface{}), 3, new(Dlink)}
	lru.set("liukai", "hubei")
	lru.set("youyi", "fujian")
	fmt.Println(lru.get("liukai"))
	fmt.Println(lru.get("youyi"))
	lru.set("dongye", "shandong")
	fmt.Println(lru.get("liukai"))
	fmt.Println(lru.get("youyi"))
	lru.set("minghai", "dongbei")
	lru.printLink()
}

type Data struct {
	key   string
	value interface{}
}

type Node struct {
	pre  *Node
	data *Data
	next *Node
}

type Dlink struct {
	head *Node
	tail *Node
}

//利用双向链表和hash表实现lru算法
type LRU struct {
	hmap map[string]interface{}
	max  int //缓存的最大容量
	dl   *Dlink
}

func (lru *LRU) printLink() {
	head := lru.dl.head
	for head != nil {
		fmt.Printf("%v:%v|", head.data.key, head.data.value)
		head = head.next
	}
	fmt.Println()
}

func (lru *LRU) set(key string, value interface{}) bool {
	if lru == nil || lru.max <= 0 {
		return false
	}
	//先判断key是否存在，如果存在则将对应双链表中的node移动至链表头部
	if _, ok := lru.hmap[key]; ok {
		lru.dl.moveToHead(key)
	} else {
		//不存在，则判断缓存容量是否已满
		if len(lru.hmap) >= lru.max {
			//容量已满，则删除双链表尾部node，并将新node插入链表首部，同时需要更新hmap
			delKey := lru.dl.removeTail()
			delete(lru.hmap, delKey)
			lru.dl.insertHead(key, value)
			lru.hmap[key] = value
		} else {
			//容量未满，则直接将新node插入链表首部，并更新hmap
			lru.dl.insertHead(key, value)
			lru.hmap[key] = value
		}
	}
	return true
}

func (lru *LRU) get(key string) interface{} {
	if v, ok := lru.hmap[key]; ok {
		lru.dl.moveToHead(key)
		return v
	}
	return nil
}

func (dl *Dlink) moveToHead(key string) {
	head := dl.head
	var target *Node
	for head != nil {
		if head.data.key == key {
			target = head
			break
		}
		head = head.next
	}
	//如果target.pre为nil，说明target就是头节点，不需要移动
	if target.pre != nil {
		pre := target.pre
		next := target.next
		pre.next = next
		if next != nil {
			next.pre = pre
		} else {
			//next为nil，说明target是尾节点，此时需要更新dl的tail
			dl.tail = pre
		}
		target.pre = nil
		target.next = dl.head
		dl.head.pre = target
		dl.head = target
	}
}

func (dl *Dlink) insertHead(key string, value interface{}) {
	newHead := new(Node)
	newHead.data = &Data{key, value}
	if dl.head == nil {
		dl.head = newHead
		dl.tail = newHead
	} else {
		newHead.next = dl.head
		dl.head.pre = newHead
		dl.head = newHead
	}
}

func (dl *Dlink) removeTail() string {
	if dl.tail == nil {
		return ""
	}
	newTail := dl.tail.pre
	delKey := dl.tail.data.key
	if newTail == nil {
		//只有一个节点
		dl.head = nil
		dl.tail = nil
	} else {
		newTail.next = nil
		dl.tail = newTail
	}
	return delKey
}
