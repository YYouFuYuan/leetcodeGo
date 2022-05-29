package main

type LRUCache struct {
	cap     int
	dList   *DoubleList
	mapNode map[int]*Node
}

//构建对象
func Constructor(capacity int) LRUCache {
	lru := &LRUCache{
		cap:     capacity,
		dList:   NewDoubleList(),
		mapNode: make(map[int]*Node),
	}
	return *lru
}

//更新节点位置
func (this *LRUCache) Update(node *Node) {
	this.dList.deleteNode(node)
	this.dList.addNodeLast(node)
}

//获得节点
func (this *LRUCache) Get(key int) int {
	node, ok := this.mapNode[key]
	if ok {
		this.Update(node)
		return node.val
	} else {
		return -1
	}

}

//放入节点
func (this *LRUCache) Put(key int, value int) {
	node, ok := this.mapNode[key]
	if ok {
		node.val = value
		this.Update(node)
	} else {
		//如果已经到达个数，删除队头
		if len(this.mapNode) >= this.cap {
			deleteNode := this.dList.head.next
			this.dList.deleteNode(deleteNode)
			delete(this.mapNode, deleteNode.key)
		}
		//加入节点
		node = NewNode(key, value)
		this.mapNode[key] = node
		this.dList.addNodeLast(node)
	}
}

type DoubleList struct {
	head *Node
	tail *Node
}

//生成一个双向链表
func NewDoubleList() *DoubleList {
	dList := &DoubleList{
		head: NewNode(0, 0),
		tail: NewNode(0, 0),
	}
	dList.head.next = dList.tail
	dList.tail.prev = dList.head
	return dList
}

//删除节点
func (this *DoubleList) deleteNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

//加入节点，放在最后的位置
func (this *DoubleList) addNodeLast(node *Node) {
	node.prev = this.tail.prev
	this.tail.prev.next = node
	node.next = this.tail
	this.tail.prev = node

}

type Node struct {
	val  int
	key  int
	next *Node
	prev *Node
}

//生成一个节点
func NewNode(key, val int) *Node {
	node := &Node{val, key, nil, nil}
	return node
}
